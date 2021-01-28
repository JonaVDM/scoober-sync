package sync

// Log hold a log, thanks vscode for letting me type this out
type Log struct {
	Added   []string
	Deleted []string
	Updated []string
}

func newLog() *Log {
	return &Log{
		Added:   make([]string, 0),
		Deleted: make([]string, 0),
		Updated: make([]string, 0),
	}
}
