import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import Typography from '@material-ui/core/Typography'
import TableRow from '@material-ui/core/TableRow'
import TableCell from '@material-ui/core/TableCell'
import Card from '@material-ui/core/Card'
import Link from 'components/Link'
import TimeAgo from 'components/TimeAgo'
import ReactStaticLinkComponent from '../ReactStaticLinkComponent'
import classNames from 'classnames'
import Button from 'components/Button'

const styles = theme => {
  return {
    jobRunsCard: {
      overflow: 'auto'
    },
    idCell: {
      width: '40%'
    },
    stampCell: {
      width: '40%'
    },
    statusCell: {
      textAlign: 'center'
    },
    runDetails: {
      paddingTop: theme.spacing.unit * 2,
      paddingBottom: theme.spacing.unit * 2
    },
    status: {
      paddingLeft: theme.spacing.unit * 1.55,
      paddingRight: theme.spacing.unit * 1.55,
      paddingTop: theme.spacing.unit / 2.1,
      paddingBottom: theme.spacing.unit / 2.1,
      borderRadius: theme.spacing.unit * 2,
      width: 'fit-content',
      display: 'inline-block'
    },
    failed: {
      backgroundColor: '#e9faf2',
      color: '#ff6587'
    },
    pending: {
      backgroundColor: '#fef7e5',
      color: '#fecb4c'
    },
    complete: {
      backgroundColro: '#e9faf2',
      color: '#4ed495'
    }
  }
}

const statusText = status => {
  if (status === 'in_progress') {
    return 'Pending'
  }
  if (status === 'pending_confirmations') {
    return 'Pending Confirmations'
  }
  if (status === 'pending_sleep') {
    return 'Pending Sleep'
  }
  if (status === 'pending_connection') {
    return 'Pending Connection'
  }
  if (status === 'pending_bridge') {
    return 'Pending Bridge'
  }
  if (status === 'error') {
    return 'Failed'
  }
  if (status === 'completed') {
    return 'Complete'
  }
}

const renderRuns = (runs, classes) => {
  if (runs && runs.length === 0) {
    return (
      <TableRow>
        <TableCell colSpan={5}>The job hasn’t run yet</TableCell>
      </TableRow>
    )
  } else if (runs) {
    return runs.map(r => (
      <TableRow key={r.id}>
        <TableCell className={classes.idCell} scope="row">
          <div className={classes.runDetails}>
            <Link to={`/jobs/${r.jobId}/runs/id/${r.id}`}>
              <Typography variant="h5" color="primary" component="span">
                {r.id}
              </Typography>
            </Link>
          </div>
        </TableCell>
        <TableCell className={classes.stampCell}>
          <Typography variant="body1" color="textSecondary">
            Created <TimeAgo>{r.createdAt}</TimeAgo>
          </Typography>
        </TableCell>
        <TableCell className={classes.statusCell} scope="row">
          <Typography
            variant="body1"
            className={classNames(
              classes.status,
              r.status.startsWith('pending')
                ? classes['pending']
                : classes[statusText(r.status).toLowerCase()]
            )}
          >
            {statusText(r.status)}
          </Typography>
        </TableCell>
      </TableRow>
    ))
  }

  return (
    <TableRow>
      <TableCell colSpan={5}>...</TableCell>
    </TableRow>
  )
}

const List = ({ jobSpecId, jobRuns, runs, showJobRunsCount, classes }) => (
  <Card className={classes.jobRunsCard}>
    <Table>
      <TableBody>
        {renderRuns(runs, classes)}
        {jobRuns && jobRuns.length > showJobRunsCount && (
          <TableRow>
            <TableCell>
              <div className={classes.runDetails}>
                <Button
                  to={`/jobs/${jobSpecId}/runs`}
                  component={ReactStaticLinkComponent}
                >
                  View More
                </Button>
              </div>
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  </Card>
)

List.propTypes = {
  jobSpecId: PropTypes.string.isRequired,
  runs: PropTypes.array.isRequired,
  jobRuns: PropTypes.any.isRequired,
  showJobRunsCount: PropTypes.any.isRequired
}

export default withStyles(styles)(List)
