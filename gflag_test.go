package gflag

import (
	"github.com/rfaulhaber/gargs"
	"os"
	"testing"
)

// TODO: rename these tests to something better!

func TestBool(t *testing.T) {
	resetArgs()
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
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
	resetArgs()
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
	resetArgs()
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

func TestBool4(t *testing.T) {
	resetArgs()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--testflag=false"}

	boolValue := Bool("t", "testflag", "set when testing", true)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if *boolValue {
		t.Error("reference was not set", "reference", *boolValue)
	}
}

func TestBool5(t *testing.T) {
	resetArgs()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "-t=false"}

	boolValue := Bool("t", "testflag", "set when testing", true)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if *boolValue {
		t.Error("reference was not set", "reference", *boolValue)
	}
}

func TestInt(t *testing.T) {
	resetArgs()
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
	resetArgs()
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
	resetArgs()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "-t=22"}

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
	resetArgs()
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

func TestString(t *testing.T) {
	resetArgs()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	expectedValue := "VALUE"
	defaultValue := "defaultValue"

	os.Args = []string{"cmd", "-t", "VALUE"}

	stringValue := String("t", "testflag", "set when testing", defaultValue)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if *stringValue != expectedValue {
		t.Error("reference was not set", "stringValue", *stringValue)
	}
}

func TestString2(t *testing.T) {
	resetArgs()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	expectedValue := "VALUE"
	defaultValue := "defaultValue"

	os.Args = []string{"cmd", "-t=VALUE"}

	stringValue := String("t", "testflag", "set when testing", defaultValue)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if *stringValue != expectedValue {
		t.Error("reference was not set", "intValue", *stringValue)
	}
}

func TestString3(t *testing.T) {
	resetArgs()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	expectedValue := "VALUE"
	defaultValue := "defaultValue"

	os.Args = []string{"cmd", "--testflag=VALUE"}

	stringValue := String("t", "testflag", "set when testing", defaultValue)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if *stringValue != expectedValue {
		t.Error("reference was not set", "intValue", *stringValue)
	}
}

func TestString4(t *testing.T) {
	resetArgs()
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	expectedValue := "VALUE"
	defaultValue := "defaultValue"

	os.Args = []string{"cmd", "--testflag", "VALUE"}

	stringValue := String("t", "testflag", "set when testing", defaultValue)

	_, ok := FlagMap["t"]

	if !ok {
		t.Error("Key was not added to flagmap")
	}

	Parse()

	if *stringValue != expectedValue {
		t.Error("reference was not set", "intValue", *stringValue)
	}
}

func TestMain(m *testing.M) {
	// TODO: add Parse() here
	retCode := m.Run()
	os.Exit(retCode)
}

func resetArgs() {
	FlagMap = make(map[string]*Flag)
	gargs.Args = make([]string, 0)
	gargs.ArgsMap = make(map[string]string)
	gargs.FlagMap = make(map[string]gargs.FlagType)
}
