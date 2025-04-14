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
		return "To do"
	case InProgress:
		return "In Progress"
	case Done:
		return "Done"
	default:
		return "Unknown"
	}
}

func IsValidTaskStatus(status TaskStatus) bool {
	switch status {
	case Todo, InProgress, Done:
		return true
	default:
		return false
	}
}