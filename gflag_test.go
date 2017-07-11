package gflag

import (
	"os"
	"testing"
)

// TODO: rename these tests to something better!

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
		t.Error("reference was not set", "reference", *boolValue)
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
		t.Error("reference was not set", "reference", *boolValue)
	}
}

func TestBool3(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--testflag"}

	boolValue := Bool("t", "testflag", "set when testing", false)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if !*boolValue {
		t.Error("reference was not set", "reference", *boolValue)
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

	if *intValue != 22 {
		t.Error("reference was not set", "intValue", *intValue)
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

	if *intValue != 22 {
		t.Error("reference was not set", "intValue", *intValue)
	}
}

func TestInt3(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--t=22"}

	intValue := Int("t", "", "set when testing", 9)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if *intValue != 22 {
		t.Error("reference was not set", "intValue", *intValue)
	}
}

func TestInt4(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "-t", "22"}

	intValue := Int("t", "testflag", "set when testing", 9)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if *intValue != 22 {
		t.Error("reference was not set", "intValue", *intValue)
	}
}

func TestMain(m *testing.M) {
	// TODO: add Parse() here
	retCode := m.Run()
	os.Exit(retCode)
}
