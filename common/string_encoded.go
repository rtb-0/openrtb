package common

import (
	"bytes"
	"encoding/json"
	"reflect"
)

// MarshalString encodes s as a JSON string without escaping <, >, or &.
func MarshalString(s string) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(s); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}

// StringEncoded carries a typed value that OpenRTB stores as a JSON string.
// Object values such as a native response are wrapped into that string.
// A value whose own JSON encoding is already a string is written as-is.
type StringEncoded[A any] struct {
	Value A
}

// Encode wraps v so it can be stored in a string-encoded OpenRTB field.
func Encode[A any](v A) StringEncoded[A] {
	return StringEncoded[A]{Value: v}
}

// IsZero reports whether Value is the zero value so encoding/json omitempty can drop the field.
func (s StringEncoded[A]) IsZero() bool {
	return reflect.ValueOf(&s.Value).Elem().IsZero()
}

// UnmarshalJSON passes the raw JSON token through to A.
// A JSON string reaches A's UnmarshalJSON when A implements it, and fills a plain string otherwise.
func (s *StringEncoded[A]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		var zero A
		s.Value = zero
		return nil
	}
	return json.Unmarshal(data, &s.Value)
}

// MarshalJSON writes Value as the JSON token of a string field.
// An encoding that already starts with a quote is used unchanged.
// Any other JSON value is wrapped in a JSON string.
func (s StringEncoded[A]) MarshalJSON() ([]byte, error) {
	if str, ok := any(s.Value).(string); ok {
		return MarshalString(str)
	}
	raw, err := json.Marshal(s.Value)
	if err != nil {
		return nil, err
	}
	if len(raw) == 0 || raw[0] == '"' {
		return raw, nil
	}
	return MarshalString(string(raw))
}

// UnwrapNativeRoot accepts a JSON object or a JSON string that contains one.
// A Native 1.0 document wrapped as {"native": {...}} is unwrapped to that object.
func UnwrapNativeRoot(data []byte) ([]byte, error) {
	raw, err := UnwrapJSONString(data)
	if err != nil {
		return nil, err
	}
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 || raw[0] != '{' {
		return raw, nil
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	tok, err := dec.Token()
	if err != nil || tok != json.Delim('{') {
		return raw, nil
	}
	var native json.RawMessage
	keys := 0
	for dec.More() {
		keyTok, err := dec.Token()
		if err != nil {
			return raw, nil
		}
		var val json.RawMessage
		if err := dec.Decode(&val); err != nil {
			return raw, nil
		}
		keys++
		if key, ok := keyTok.(string); ok && key == "native" {
			native = val
		}
	}
	if keys == 1 && len(native) > 0 {
		return native, nil
	}
	return raw, nil
}

// UnwrapJSONString returns the decoded text when data is a JSON string.
// Any other token, including a JSON object, is returned unchanged.
func UnwrapJSONString(data []byte) ([]byte, error) {
	if len(data) == 0 || data[0] != '"' {
		return data, nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return []byte(s), nil
}
