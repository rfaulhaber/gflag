package gflag

import (
	"testing"
	"fmt"
)

func TestBool(t *testing.T) {
	Bool("t", "testflag", "set when testing", false)

	expected, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}
}
