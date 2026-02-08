package llm

type ErrKind string

const (
	ErrTransient ErrKind = "transient"
	ErrInvalid   ErrKind = "invalid"
	ErrPolicy    ErrKind = "policy"
	ErrUnknown   ErrKind = "unknown"
)

type Error struct {
	Kind  ErrKind
	Msg   string
	Cause error
}

func (e *Error) Error() string { return string(e.Kind) + ":" + e.Msg }
func (e *Error) Unwrap() error { return e.Cause }
