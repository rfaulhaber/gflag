package gflag

import (
	"os"
	"strconv"
)

type Value interface {
	String() string
	Set(string) error
}

// specifies boolean flags
// by virtue of placing them they are set to true
type BoolVal struct {
	value bool
}

func (b BoolVal) String() string {
	if b.value {
		return "true"
	} else {
		return "false"
	}
}

func (b *BoolVal) Set(value string) error {
	parsed := strconv.ParseBool(value)
}

type StringVal struct {
	value string
}

type Flag struct {
	ShortName string
	LongName string
	Usage string
	Value Value
}

func Bool(shortName string, longName string, usage string) *Flag {
	return &Flag{shortName, longName, usage, BoolVal{true}}
}

type FlagSet struct {
	Usage func()

	flags map[string]*Flag
}

func Parse() {

}
