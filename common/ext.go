package common

import (
	"encoding/json"

	"github.com/rtb-0/openrtb/errs"
)

// Ext is raw JSON for an exchange-specific extension.
// The backing array holds no pointers, so the garbage collector does not scan it.
//
// Ext is its own type so it can have methods. MarshalJSON and UnmarshalJSON
// forward to json.RawMessage, so a field still encodes as raw JSON.
type Ext json.RawMessage

// MarshalJSON returns the raw JSON. Nil encodes as null.
func (e Ext) MarshalJSON() ([]byte, error) {
	return json.RawMessage(e).MarshalJSON()
}

var (
	// ErrNilPointer is returned when UnmarshalJSON is called on a nil Ext.
	ErrNilPointer = errs.Err("unmarshal on nil pointer")
	// ErrExt is a failure to decode an extension into a typed value.
	ErrExt = errs.Err("ext")
)

// UnmarshalJSON copies data into e.
func (e *Ext) UnmarshalJSON(data []byte) error {
	if e == nil {
		return ErrNilPointer.WithMessage("Ext")
	}
	*e = append((*e)[:0], data...)
	return nil
}

// As unmarshals the extension into T.
// An empty extension returns the zero value.
func (e Ext) As[T any]() (T, error) {
	var v T
	if len(e) == 0 {
		return v, nil
	}
	if err := json.Unmarshal([]byte(e), &v); err != nil {
		return v, ErrExt.Wrap(err)
	}
	return v, nil
}

// MustAs unmarshals the extension into T.
// It panics when the JSON is invalid.
func (e Ext) MustAs[T any]() T {
	v, err := e.As[T]()
	if err != nil {
		panic(err)
	}
	return v
}

// AsOr unmarshals the extension into T.
// An empty or invalid extension returns fallback.
func (e Ext) AsOr[T any](fallback T) T {
	if len(e) == 0 {
		return fallback
	}
	v, err := e.As[T]()
	if err != nil {
		return fallback
	}
	return v
}
