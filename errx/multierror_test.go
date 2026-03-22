package errx

import (
	"errors"
	"strings"
	"testing"
)

func TestMultiError(t *testing.T) {
	m := &MultiError{}

	if m.HasErrors() {
		t.Error("expected new MultiError to be empty")
	}

	if m.AsError() != nil {
		t.Error("expected AsError to return nil when empty")
	}

	m.Append(nil) // Should ignore nils natively

	err1 := errors.New("first log failure")
	err2 := errors.New("secondary mathematical log string")

	m.Append(err1)
	m.Append(err2)

	if !m.HasErrors() {
		t.Error("expected boolean to be true")
	}

	err := m.AsError()
	if err == nil {
		t.Error("expected structured standard error strictly non-nil")
	}

	msg := err.Error()
	if !strings.Contains(msg, "first log failure") || !strings.Contains(msg, "secondary mathematical log") {
		t.Errorf("error implicitly standard formatting failed: %s", msg)
	}
}
