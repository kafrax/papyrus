package cape

type ErrorType = uint64

const (
	ErrorServer ErrorType = 1 << 63
)

type Error struct {
	Err     error
	ErrType ErrorType
	ErrMsg  string
}

func (e *Error) Error() string {
	return e.ErrMsg
}

func ErrorsNew(msg string, t ErrorType, err error) Error {
	return Error{err, t, msg}
}
