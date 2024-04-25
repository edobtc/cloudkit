package multiplex

import "fmt"

type Behavior int

const (
	Unknown Behavior = iota
	ReadWrite
	Read
	Write
)

func (b Behavior) String() string {
	switch b {
	case ReadWrite:
		return "ReadWrite"
	case Read:
		return "Read"
	case Write:
		return "Write"
	default:
		return "Unknown"
	}
}

func BehaviorFromString(s string) (Behavior, error) {
	switch s {
	case "ReadWrite", "readwrite", "rw":
		return ReadWrite, nil
	case "Read", "read", "r":
		return Read, nil
	case "Write", "write", "w":
		return Write, nil
	case "Unknown", "unknown", "u":
		return Unknown, nil
	default:
		return 0, fmt.Errorf("invalid behavior string: %s", s)
	}
}
