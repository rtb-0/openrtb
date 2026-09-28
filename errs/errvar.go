// Package errs declares reusable errors.
// Err returns a sentinel. WithMessage and Wrap add detail that errors.Is still matches.
package errs

import (
	"fmt"
	"sync/atomic"
)

// ErrVar is a reusable error declared ahead of time.
// WithMessage, WithMessageFmt, and Wrap attach detail without changing the sentinel,
// so errors.Is still matches the original value.
type ErrVar struct {
	id   uint32
	text string
}

var errSeq atomic.Uint32

// Err declares a reusable error.
func Err(text string) ErrVar {
	return ErrVar{id: errSeq.Add(1), text: text}
}

func (v ErrVar) Error() string { return v.text }

// WithMessage returns v plus an extra explanation.
func (v ErrVar) WithMessage(msg string) *Error {
	return &Error{ErrVar: v, msg: msg}
}

// WithMessageFmt returns v plus an explanation formatted as fmt.Sprintf.
func (v ErrVar) WithMessageFmt(format string, args ...any) *Error {
	return v.WithMessage(fmt.Sprintf(format, args...))
}

// Wrap returns v with err as its cause.
// A nil cause returns nil.
func (v ErrVar) Wrap(err error) *Error {
	if err == nil {
		return nil
	}
	return &Error{ErrVar: v, cause: err}
}
