package gflag

import (
	"github.com/rfaulhaber/gargs"
	"os"
	"testing"
)

func TestBool(t *testing.T) {
	boolValue := Bool("t", "testflag", "set when testing", false)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if !*boolValue {
		t.Error("reference was not set")
	}
}

func TestMain(m *testing.M) {
	initialize()
	// TODO: add Parse() here
	retCode := m.Run()
	os.Exit(retCode)
}

func initialize() {
	gargs.Args = []string{"testflag"}
}
