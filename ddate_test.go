package ddate

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	want := "Sweetmorn, Chaos 1, 3136 YOLD"
	unixStart := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

	msg := Format(unixStart)
	if want != msg {
		t.Fatalf(`1.1.1970 = %q, want %q`, msg, want)
	}
}

func TestIsHolyday(t *testing.T) {
	want := "Mungday"
	holy := time.Date(1970, 1, 5, 0, 0, 0, 0, time.UTC)

	celebrate, msg := IsHolyday(holy)

	if want != msg || !celebrate {
		t.Fatalf(`5.1.1970 must be a %q`, want)
	}
}

func TestIsNotHolyday(t *testing.T) {
	notholy := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

	celebrate, msg := IsHolyday(notholy)

	if celebrate || msg != "" {
		t.Fatalf(`1.1.1970 was not a holyday`)
	}
}
