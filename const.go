package pgxhook

type CallerType string

const (
	CallerExec      CallerType = "exec"
	CallerQuery     CallerType = "query"
	CallerQueryRow  CallerType = "query_row"
	CallerSendBatch CallerType = "send_batch"
	CallerPing      CallerType = "ping"
	CallerClose     CallerType = "close"
	CallerPrepare   CallerType = "prepare"
	CallerBegin     CallerType = "begin"
	CallerCommit    CallerType = "commit"
	CallerRollback  CallerType = "rollback"
)
