package loader

const (
	FAILED = iota
	OK
	PAUSED
)

type Loader struct {
	state int
}

func (l *Loader) State() string {
	switch state {
	case FAILED:
		return "[" + "FAILED" + "]"
	case OK:
		return "[" + "  OK  " + "]"
	case PAUSED:
		return "[" + "PAUSED" + "]"
	}
}
