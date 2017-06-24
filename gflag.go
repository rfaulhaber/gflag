package gflag

import "io"

type FlagSet struct {
	Usage func()

	shortName string
	longName string
	args []string
	output io.Writer
}

type Flag struct {
	ShortName string
	LongName string
	Usage string
	DefaultValue string
}
