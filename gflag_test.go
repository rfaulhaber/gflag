package gflag

import (
	"os"
	"testing"
)

func TestBool(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "-t"}

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
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--testflag"}

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

func TestInt(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--testflag=22"}
	intValue := Int("", "testflag", "set when testing", 9)

	_, ok := FlagMap["testflag"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if *intValue != 22 || *intValue == 9 {
		t.Error("reference was not set")
	}
}

func TestInt2(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--testflag", "22"}
	intValue := Int("", "testflag", "set when testing", 9)

	_, ok := FlagMap["testflag"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if *intValue != 22 || *intValue == 9 {
		t.Error("reference was not set")
	}
}

func TestMain(m *testing.M) {
	// TODO: add Parse() here
	retCode := m.Run()
	os.Exit(retCode)
}
