package flakeytests

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

type mockReporter struct {
	report *Report
}

func (m *mockReporter) Report(_ context.Context, report *Report) error {
	m.report = report
	return nil
}

func newMockReporter() *mockReporter {
	return &mockReporter{report: NewReport()}
}

func TestParser(t *testing.T) {
	tableTests := []struct {
		name    string
		output  string
		asserts func(ts map[string]map[string]int)
		err     error
	}{
		{
			name:   "testParser- 01",
			output: `{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`,
			asserts: func(ts map[string]map[string]int) {
				assert.Len(t, ts, 1)
				assert.Len(t, ts["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"], 1)
				assert.Equal(t, ts["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"]["TestLink"], 1)
			},
			err: nil,
		},
		{
			name: "skips non JSON",
			output: `Failed tests and panics:
-------
{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}
`,
			asserts: func(ts map[string]map[string]int) {
				assert.Len(t, ts, 1)
				assert.Len(t, ts["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"], 1)
				assert.Equal(t, ts["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"]["TestLink"], 1)
			},
			err: nil,
		},
		{
			name: "Panic due to logging",
			output: `
{"Time":"2023-09-07T16:01:40.649849+01:00","Action":"output","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestAssets_LinkScanValue","Output":"panic: foo\n"}
`,
			asserts: func(ts map[string]map[string]int) {
				assert.Len(t, ts, 1)
				assert.Len(t, ts["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"], 1)
				assert.Equal(t, ts["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"]["TestAssets_LinkScanValue"], 1)
			},
			err: nil,
		},
		{
			name: "Successful output",
			output: `
{"Time":"2023-09-07T16:22:52.556853+01:00","Action":"start","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"}
{"Time":"2023-09-07T16:22:52.762353+01:00","Action":"run","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestAssets_NewLinkAndString"}
{"Time":"2023-09-07T16:22:52.762456+01:00","Action":"output","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestAssets_NewLinkAndString","Output":"=== RUN   TestAssets_NewLinkAndString\n"}
{"Time":"2023-09-07T16:22:52.76249+01:00","Action":"output","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestAssets_NewLinkAndString","Output":"=== PAUSE TestAssets_NewLinkAndString\n"}
{"Time":"2023-09-07T16:22:52.7625+01:00","Action":"pause","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestAssets_NewLinkAndString"}
{"Time":"2023-09-07T16:22:52.762511+01:00","Action":"cont","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestAssets_NewLinkAndString"}
{"Time":"2023-09-07T16:22:52.762528+01:00","Action":"output","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestAssets_NewLinkAndString","Output":"=== CONT  TestAssets_NewLinkAndString\n"}
{"Time":"2023-09-07T16:22:52.762546+01:00","Action":"output","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestAssets_NewLinkAndString","Output":"--- PASS: TestAssets_NewLinkAndString (0.00s)\n"}
{"Time":"2023-09-07T16:22:52.762557+01:00","Action":"pass","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestAssets_NewLinkAndString","Elapsed":0}
{"Time":"2023-09-07T16:22:52.762566+01:00","Action":"output","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Output":"PASS\n"}
{"Time":"2023-09-07T16:22:52.762955+01:00","Action":"output","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Output":"ok  \tgithub.com/smartcontractkit/chainlink/v2/core/assets\t0.206s\n"}
{"Time":"2023-09-07T16:22:52.765598+01:00","Action":"pass","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Elapsed":0.209}
`,
			asserts: func(ts map[string]map[string]int) {
				assert.Len(t, ts, 0)
			},
			err: nil,
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.output)
			pr, err := parseOutput(r)
			require.Equal(t, tt.err, err)
			ts := pr.tests
			tt.asserts(ts)
		})
	}
}

type testAdapter func(string, []string, io.Writer) error

func (t testAdapter) test(pkg string, tests []string, out io.Writer) error {
	return t(pkg, tests, out)
}

func TestRunner(t *testing.T) {
	tableTests := []struct {
		name          string
		initialOutput []string
		outputs       []string
		asserts       func(m *mockReporter, j int)
	}{
		{
			name:          "testRunner with flake",
			initialOutput: []string{`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`},
			outputs: []string{
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`,
				``,
			},
			asserts: func(m *mockReporter, j int) {
				assert.Len(t, m.report.tests, 1)
				_, ok := m.report.tests["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"]["TestLink"]
				assert.True(t, ok)
			},
		},
		{
			name: "testRunner with failed package",
			initialOutput: []string{`
{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}
{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Elapsed":0}
`},
			outputs: []string{`
{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}
{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Elapsed":0}
`,
				``,
			},
			asserts: func(m *mockReporter, j int) {
				assert.Len(t, m.report.tests, 1)
				_, ok := m.report.tests["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"]["TestLink"]
				assert.True(t, ok)
			},
		},
		{
			name:          "testRunner rerun successful",
			initialOutput: []string{`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`},
			outputs: []string{
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`,
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"pass","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`,
			},
			asserts: func(m *mockReporter, j int) {
				_, ok := m.report.tests["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"]["TestLink"]
				assert.True(t, ok)
			},
		},
		{
			name:          "testRunner rerun fails with non zero exit code",
			initialOutput: []string{`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`},
			outputs: []string{
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`,
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"pass","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`,
			},
			asserts: func(m *mockReporter, j int) {
				_, ok := m.report.tests["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"]["TestLink"]
				assert.True(t, ok)
			},
		},
		{
			name: "testRunner rerun with non zero exit code does not stop command",
			initialOutput: []string{
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`,
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/services/vrf/v2","Test":"TestMaybeReservedLinkV2","Elapsed":0}`,
			},
			outputs: []string{
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`,
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"pass","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`,
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/services/vrf/v2","Test":"TestMaybeReservedLinkV2","Elapsed":0}`,
				`{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/services/vrf/v2","Test":"TestMaybeReservedLinkV2","Elapsed":0}`,
			},
			asserts: func(m *mockReporter, j int) {
				_, ok := m.report.tests["github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"]["TestLink"]
				assert.True(t, ok)
				assert.Equal(t, 4, j)
			},
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			var initOutput []io.Reader
			for i := range tt.initialOutput {
				initOutput = append(initOutput, strings.NewReader(tt.initialOutput[i]))
			}
			m := newMockReporter()
			j := 0
			r := &Runner{
				numReruns: 2,
				readers:   initOutput,
				testCommand: testAdapter(func(pkg string, testNames []string, w io.Writer) error {
					_, err := w.Write([]byte(tt.outputs[j]))
					j++
					return err
				}),
				parse:    parseOutput,
				reporter: m,
			}
			// This will report a flake since we've mocked the rerun
			// to only report one failure (not two as expected).
			err := r.Run(tests.Context(t))
			require.NoError(t, err)
			tt.asserts(m, j)
		})
	}
}

func TestRunner_AllFailures(t *testing.T) {
	tableTests := []struct {
		name        string
		output      string
		rerunOutput string
		asserts     func(m *mockReporter)
	}{
		{
			name:   "all failures",
			output: `{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}`,
			rerunOutput: `
{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}
{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets","Test":"TestLink","Elapsed":0}
`,
			asserts: func(m *mockReporter) {
				assert.Len(t, m.report.tests, 0)
			},
		},
		{
			name:        "root level test",
			output:      `{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/","Test":"TestConfigDocs","Elapsed":0}`,
			rerunOutput: ``,
			asserts: func(m *mockReporter) {
				_, ok := m.report.tests["github.com/smartcontractkit/chainlink/v2/"]["TestConfigDocs"]
				assert.True(t, ok)
			},
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMockReporter()
			r := &Runner{
				numReruns: 2,
				readers:   []io.Reader{strings.NewReader(tt.output)},
				testCommand: testAdapter(func(pkg string, testNames []string, w io.Writer) error {
					_, err := w.Write([]byte(tt.rerunOutput))
					return err
				}),
				parse:    parseOutput,
				reporter: m,
			}

			err := r.Run(tests.Context(t))
			require.NoError(t, err)
			tt.asserts(m)
		})
	}
}

// Used for integration tests
func TestSkippedForTests_Subtests(t *testing.T) {
	if os.Getenv("FLAKEY_TEST_RUNNER_RUN_FIXTURE_TEST") != "1" {
		t.Skip()
	}

	t.Run("1: should fail", func(t *testing.T) {
		assert.False(t, true)
	})

	t.Run("2: should fail", func(t *testing.T) {
		assert.False(t, true)
	})
}

// Used for integration tests
func TestSkippedForTests(t *testing.T) {
	if os.Getenv("FLAKEY_TEST_RUNNER_RUN_FIXTURE_TEST") != "1" {
		t.Skip()
	}

	go func() {
		panic("skipped test")
	}()
}

// Used for integration tests
func TestSkippedForTests_Success(t *testing.T) {
	if os.Getenv("FLAKEY_TEST_RUNNER_RUN_FIXTURE_TEST") != "1" {
		t.Skip()
	}

	assert.True(t, true)
}

func TestIntegration_DealsWithSubtests(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	tableTests := []struct {
		name    string
		output  string
		asserts func(m *mockReporter)
	}{
		{
			name: "deals with subtests",
			output: `
{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/tools/flakeytests/","Test":"TestSkippedForTests_Subtests/1:_should_fail","Elapsed":0}
{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/tools/flakeytests/","Test":"TestSkippedForTests_Subtests","Elapsed":0}
{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/tools/flakeytests/","Test":"TestSkippedForTests_Subtests/2:_should_fail","Elapsed":0}
`,
			asserts: func(m *mockReporter) {
				expectedTests := map[string]map[string]int{
					"github.com/smartcontractkit/chainlink/v2/tools/flakeytests/": {
						"TestSkippedForTests_Subtests/1:_should_fail": 1,
						"TestSkippedForTests_Subtests/2:_should_fail": 1,
					},
				}
				assert.Equal(t, expectedTests, m.report.tests)
			},
		},
		{
			name:   "parse panics",
			output: `{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/tools/flakeytests/","Test":"TestSkippedForTests","Elapsed":0}`,
			asserts: func(m *mockReporter) {
				_, ok := m.report.tests["github.com/smartcontractkit/chainlink/v2/tools/flakeytests"]["TestSkippedForTests"]
				assert.False(t, ok)
			},
		},
		{
			name:   "test integration",
			output: `{"Time":"2023-09-07T15:39:46.378315+01:00","Action":"fail","Package":"github.com/smartcontractkit/chainlink/v2/tools/flakeytests/","Test":"TestSkippedForTests_Success","Elapsed":0}`,
			asserts: func(m *mockReporter) {
				_, ok := m.report.tests["github.com/smartcontractkit/chainlink/v2/tools/flakeytests"]["TestSkippedForTests_Success"]
				assert.False(t, ok)
			},
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMockReporter()
			tc := &testCommand{
				repo:    "github.com/smartcontractkit/chainlink/v2/tools/flakeytests",
				command: "../bin/go_core_tests",
				overrides: func(cmd *exec.Cmd) {
					cmd.Env = append(cmd.Env, "FLAKEY_TESTRUNNER_RUN_FIXTURE_TEST=1")
					cmd.Stdout = io.Discard
					cmd.Stderr = io.Discard
				},
			}
			r := &Runner{
				numReruns:   2,
				readers:     []io.Reader{strings.NewReader(tt.output)},
				testCommand: tc,
				parse:       parseOutput,
				reporter:    m,
			}

			err := r.Run(tests.Context(t))
			require.NoError(t, err)
			tt.asserts(m)
		})
	}
}
