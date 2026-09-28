package errs

import "fmt"

// Error is an ErrVar plus an optional explanation and a wrapped cause.
type Error struct {
	ErrVar ErrVar
	msg    string
	cause  error
}

func (e *Error) Error() string {
	s := e.ErrVar.text
	if e.msg != "" {
		s += ": " + e.msg
	}
	if e.cause != nil {
		s += ": " + e.cause.Error()
	}
	return s
}

// Is reports whether target is the sentinel stored in e.
func (e *Error) Is(target error) bool {
	v, ok := target.(ErrVar)
	return ok && v == e.ErrVar
}

// Unwrap returns the wrapped cause.
func (e *Error) Unwrap() error { return e.cause }

// WithMessage sets the explanation and returns a copy.
func (e *Error) WithMessage(msg string) *Error {
	if e == nil {
		return nil
	}
	cp := *e
	cp.msg = msg
	return &cp
}

// WithMessageFmt sets the explanation from a format string and returns a copy.
func (e *Error) WithMessageFmt(format string, args ...any) *Error {
	return e.WithMessage(fmt.Sprintf(format, args...))
}

// Wrap sets the cause and returns a copy.
// A nil cause returns nil.
func (e *Error) Wrap(err error) *Error {
	if e == nil || err == nil {
		return nil
	}
	cp := *e
	cp.cause = err
	return &cp
}
