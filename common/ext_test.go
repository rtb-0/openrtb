package common

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestExtJSONUnchanged(t *testing.T) {
	in := []byte(`{"ext":{"seat":"a","n":1}}`)
	var got struct {
		Ext Ext `json:"ext"`
	}
	if err := json.Unmarshal(in, &got); err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != string(in) {
		t.Fatalf("ext json = %s", out)
	}
}

func TestAs(t *testing.T) {
	type seat struct {
		Seat string `json:"seat"`
		N    int    `json:"n"`
	}
	got, err := Ext(`{"seat":"a","n":1}`).As[seat]()
	if err != nil {
		t.Fatal(err)
	}
	if got.Seat != "a" || got.N != 1 {
		t.Fatalf("as = %+v", got)
	}

	var none Ext
	empty, err := none.As[seat]()
	if err != nil {
		t.Fatal(err)
	}
	if empty != (seat{}) {
		t.Fatalf("empty = %+v", empty)
	}

	if got := Ext(`{"seat":"a","n":1}`).MustAs[seat](); got != (seat{Seat: "a", N: 1}) {
		t.Fatalf("must = %+v", got)
	}
	fallback := seat{Seat: "fallback"}
	if got := Ext(nil).AsOr(fallback); got != fallback {
		t.Fatalf("empty or = %+v", got)
	}
	if got := Ext(`{`).AsOr(fallback); got != fallback {
		t.Fatalf("bad or = %+v", got)
	}
	if got := Ext(`{"seat":"a"}`).AsOr(fallback); got.Seat != "a" {
		t.Fatalf("or = %+v", got)
	}
}

func TestMustAsPanics(t *testing.T) {
	defer func() {
		rec := recover()
		err, ok := rec.(error)
		if !ok || !errors.Is(err, ErrExt) {
			t.Fatalf("panic = %v", rec)
		}
	}()
	Ext(`{`).MustAs[struct{}]()
}
