package common

import (
	"encoding/json"
	"testing"
)

type objectMarkup struct {
	Name string `json:"name"`
}

func (m *objectMarkup) UnmarshalJSON(data []byte) error {
	raw, err := UnwrapJSONString(data)
	if err != nil {
		return err
	}
	type alias objectMarkup
	var a alias
	if err := json.Unmarshal(raw, &a); err != nil {
		return err
	}
	*m = objectMarkup(a)
	return nil
}

func TestStringEncodedObjectRoundTrip(t *testing.T) {
	in := []byte(`{"adm":"{\"name\":\"banner\"}"}`)
	var got struct {
		Adm StringEncoded[objectMarkup] `json:"adm"`
	}
	if err := json.Unmarshal(in, &got); err != nil {
		t.Fatal(err)
	}
	if got.Adm.Value.Name != "banner" {
		t.Fatalf("name = %q", got.Adm.Value.Name)
	}
	out, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `{"adm":"{\"name\":\"banner\"}"}` {
		t.Fatalf("marshal = %s", out)
	}
}

func TestStringEncodedPlainString(t *testing.T) {
	var got struct {
		Adm StringEncoded[string] `json:"adm,omitempty"`
	}
	if err := json.Unmarshal([]byte(`{"adm":"<div/>"}`), &got); err != nil {
		t.Fatal(err)
	}
	if got.Adm.Value != "<div/>" {
		t.Fatalf("value = %q", got.Adm.Value)
	}
	out, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	var again struct {
		Adm StringEncoded[string] `json:"adm"`
	}
	if err := json.Unmarshal(out, &again); err != nil {
		t.Fatal(err)
	}
	if again.Adm.Value != "<div/>" {
		t.Fatalf("round trip = %q", again.Adm.Value)
	}
}

func TestStringEncodedOmitZero(t *testing.T) {
	var v struct {
		Adm StringEncoded[string] `json:"adm,omitzero"`
	}
	out, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `{}` {
		t.Fatalf("marshal = %s", out)
	}
}
