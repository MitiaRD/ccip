// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_l2_bridge_adapter

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

var MockL2BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506102fa806100206000396000f3fe6080604052600436106100345760003560e01c80632e4b1fc91461003957806338314bb21461005a57806379a35b4b1461007d575b600080fd5b34801561004557600080fd5b50604051600081526020015b60405180910390f35b34801561006657600080fd5b5061007b61007536600461017f565b50505050565b005b61009061008b36600461020d565b61009d565b6040516100519190610258565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810182905260609073ffffffffffffffffffffffffffffffffffffffff8616906323b872dd906064016020604051808303816000875af1158015610119573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061013d91906102c4565b5050604080516020810190915260008152949350505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461017a57600080fd5b919050565b6000806000806060858703121561019557600080fd5b61019e85610156565b93506101ac60208601610156565b9250604085013567ffffffffffffffff808211156101c957600080fd5b818701915087601f8301126101dd57600080fd5b8135818111156101ec57600080fd5b8860208285010111156101fe57600080fd5b95989497505060200194505050565b6000806000806080858703121561022357600080fd5b61022c85610156565b935061023a60208601610156565b925061024860408601610156565b9396929550929360600135925050565b600060208083528351808285015260005b8181101561028557858101830151858201604001528201610269565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b6000602082840312156102d657600080fd5b815180151581146102e657600080fd5b939250505056fea164736f6c6343000813000a",
}

var MockL2BridgeAdapterABI = MockL2BridgeAdapterMetaData.ABI

var MockL2BridgeAdapterBin = MockL2BridgeAdapterMetaData.Bin

func DeployMockL2BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockL2BridgeAdapter, error) {
	parsed, err := MockL2BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockL2BridgeAdapterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockL2BridgeAdapter{address: address, abi: *parsed, MockL2BridgeAdapterCaller: MockL2BridgeAdapterCaller{contract: contract}, MockL2BridgeAdapterTransactor: MockL2BridgeAdapterTransactor{contract: contract}, MockL2BridgeAdapterFilterer: MockL2BridgeAdapterFilterer{contract: contract}}, nil
}

type MockL2BridgeAdapter struct {
	address common.Address
	abi     abi.ABI
	MockL2BridgeAdapterCaller
	MockL2BridgeAdapterTransactor
	MockL2BridgeAdapterFilterer
}

type MockL2BridgeAdapterCaller struct {
	contract *bind.BoundContract
}

type MockL2BridgeAdapterTransactor struct {
	contract *bind.BoundContract
}

type MockL2BridgeAdapterFilterer struct {
	contract *bind.BoundContract
}

type MockL2BridgeAdapterSession struct {
	Contract     *MockL2BridgeAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockL2BridgeAdapterCallerSession struct {
	Contract *MockL2BridgeAdapterCaller
	CallOpts bind.CallOpts
}

type MockL2BridgeAdapterTransactorSession struct {
	Contract     *MockL2BridgeAdapterTransactor
	TransactOpts bind.TransactOpts
}

type MockL2BridgeAdapterRaw struct {
	Contract *MockL2BridgeAdapter
}

type MockL2BridgeAdapterCallerRaw struct {
	Contract *MockL2BridgeAdapterCaller
}

type MockL2BridgeAdapterTransactorRaw struct {
	Contract *MockL2BridgeAdapterTransactor
}

func NewMockL2BridgeAdapter(address common.Address, backend bind.ContractBackend) (*MockL2BridgeAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(MockL2BridgeAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockL2BridgeAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockL2BridgeAdapter{address: address, abi: abi, MockL2BridgeAdapterCaller: MockL2BridgeAdapterCaller{contract: contract}, MockL2BridgeAdapterTransactor: MockL2BridgeAdapterTransactor{contract: contract}, MockL2BridgeAdapterFilterer: MockL2BridgeAdapterFilterer{contract: contract}}, nil
}

func NewMockL2BridgeAdapterCaller(address common.Address, caller bind.ContractCaller) (*MockL2BridgeAdapterCaller, error) {
	contract, err := bindMockL2BridgeAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockL2BridgeAdapterCaller{contract: contract}, nil
}

func NewMockL2BridgeAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockL2BridgeAdapterTransactor, error) {
	contract, err := bindMockL2BridgeAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockL2BridgeAdapterTransactor{contract: contract}, nil
}

func NewMockL2BridgeAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockL2BridgeAdapterFilterer, error) {
	contract, err := bindMockL2BridgeAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockL2BridgeAdapterFilterer{contract: contract}, nil
}

func bindMockL2BridgeAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockL2BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL2BridgeAdapter.Contract.MockL2BridgeAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.MockL2BridgeAdapterTransactor.contract.Transfer(opts)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.MockL2BridgeAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL2BridgeAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.contract.Transfer(opts)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterCaller) GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockL2BridgeAdapter.contract.Call(opts, &out, "getBridgeFeeInNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _MockL2BridgeAdapter.Contract.GetBridgeFeeInNative(&_MockL2BridgeAdapter.CallOpts)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterCallerSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _MockL2BridgeAdapter.Contract.GetBridgeFeeInNative(&_MockL2BridgeAdapter.CallOpts)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20", arg0, arg1, arg2)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.FinalizeWithdrawERC20(&_MockL2BridgeAdapter.TransactOpts, arg0, arg1, arg2)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterTransactorSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.FinalizeWithdrawERC20(&_MockL2BridgeAdapter.TransactOpts, arg0, arg1, arg2)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.contract.Transact(opts, "sendERC20", localToken, arg1, arg2, amount)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterSession) SendERC20(localToken common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.SendERC20(&_MockL2BridgeAdapter.TransactOpts, localToken, arg1, arg2, amount)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterTransactorSession) SendERC20(localToken common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.SendERC20(&_MockL2BridgeAdapter.TransactOpts, localToken, arg1, arg2, amount)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapter) Address() common.Address {
	return _MockL2BridgeAdapter.address
}

type MockL2BridgeAdapterInterface interface {
	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
