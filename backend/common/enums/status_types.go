package enums

type TaskStatus int

const (
	Todo TaskStatus = iota
	InProgress
	Done
)

func (s TaskStatus) String() string {
	switch s {
	case Todo:
		return "Todo"
	case InProgress:
		return "In Progress"
	case Done:
		return "Done"
	default:
		return "Unknown"
	}
}