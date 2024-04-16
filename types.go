package pgxhook

type HookData struct {
	Query   string
	InBatch bool
	InTx    bool
	Error   error
	Caller  CallerType
	Args    []any
}
