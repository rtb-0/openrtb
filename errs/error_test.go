package errs

import (
	"errors"
	"io"
	"testing"
)

func TestErrorSentinel(t *testing.T) {
	errReq := Err("required")
	sameText := Err("required")

	got := errReq.WithMessageFmt("imp[%d].id", 0)
	if got.Error() != "required: imp[0].id" {
		t.Fatalf("message = %s", got.Error())
	}
	if !errors.Is(got, errReq) {
		t.Fatal("expected sentinel match")
	}
	if errors.Is(got, sameText) {
		t.Fatal("same text is a different sentinel")
	}

	wrapped := errReq.WithMessage("id").Wrap(io.EOF)
	if !errors.Is(wrapped, errReq) || !errors.Is(wrapped, io.EOF) {
		t.Fatalf("unwrap = %v", wrapped)
	}
	if errReq.Wrap(nil) != nil {
		t.Fatal("wrap nil")
	}
}
