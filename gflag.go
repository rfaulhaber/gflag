package gflag

import (
	"strconv"
)

type Flag struct {
	ShortName string
	LongName  string
	Usage     string
	Value     Value
}

func NewFlag(shortName string, longName string, usage string, value Value) *Flag {
	return &Flag{shortName, longName, usage, value}
}

type Value interface {
	String() string
	Set(string) error
}

type BoolValue struct {
	value bool
}

func (b BoolValue) String() string {
	return strconv.FormatBool(b.value)
}

func (b *BoolValue) Set(value string) error {
	if parsed, error := strconv.ParseBool(value); error != nil {
		return error
	} else {
		b.value = parsed
		return nil
	}
}

type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
	*p = val
	return (*boolValue)(p)
}

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

func (b *boolValue) String() string {
	return strconv.FormatBool(bool(*b))
}

type StringValue struct {
	value string
}

func (s StringValue) String() string {
	return s.value
}

func (s *StringValue) Set(value string) error {
	// if we've made it this far there shouldn't be an error!
	s.value = value
	return nil
}

type stringVal string

func newStringValue(val string, p *string) *stringVal {
	*p = val
	return (*stringVal)(p)
}

func (s *stringVal) Set(value string) error {
	*s = stringVal(value)
	return nil
}

func (s stringVal) String() string {
	return string(s)
}

type IntValue struct {
	value int
}

type intVal int

func (i *intVal) Set(val string) error {
	v, err := strconv.Atoi(val)
	*i = intVal(v)
	return err
}

func (i *intVal) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

func newIntValue(val int, p *int) *intVal {
	*p = val
	return (*intVal)(p)
}

func (i IntValue) String() string {
	return strconv.Itoa(i.value)
}

func (i *IntValue) Set(value string) error {
	if parsed, err := strconv.Atoi(value); err != nil {
		return err
	} else {
		i.value = parsed
		return nil
	}
}

var FlagMap map[string]*Flag

func init() {
	FlagMap = make(map[string]*Flag)
}

func Bool(shortName string, longName string, usage string, defaultValue bool) *bool {
	var accessor string
	if len(shortName) == 0 {
		accessor = longName
	} else {
		accessor = shortName
	}

	r := new(bool)
	flag := NewFlag(shortName, longName, usage, newBoolValue(defaultValue, r))
	setMap(accessor, flag)

	return r
}

func Int(shortName string, longName string, usage string, defaultValue int) *int {
	var accessor string
	if len(shortName) == 0 {
		accessor = longName
	} else {
		accessor = shortName
	}

	r := new(int)
	flag := NewFlag(shortName, longName, usage, newIntValue(defaultValue, r))
	setMap(accessor, flag)

	return r
}

func Parse() {
}

func setMap(accessor string, flag *Flag) {
	FlagMap[accessor] = flag
}
