package gflag

import (
	"github.com/rfaulhaber/gargs"
	"os"
	"testing"
)

func TestBool(t *testing.T) {
	gargs.Args = []string{"t"}
	boolValue := Bool("t", "", "set when testing", false)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if !*boolValue {
		t.Error("reference was not set")
	}
}

func TestBool2(t *testing.T) {
	gargs.Args = []string{"testflag"}
	boolValue := Bool("", "testflag", "set when testing", false)

	_, ok := FlagMap["testflag"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if !*boolValue {
		t.Error("reference was not set")
	}
}

func TestMain(m *testing.M) {
	// TODO: add Parse() here
	retCode := m.Run()
	os.Exit(retCode)
}
