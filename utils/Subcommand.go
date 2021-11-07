package utils

import (
	"flag"
)

type Subcommand struct {
	Fs       *flag.FlagSet
	Name     string
	Interval int
	Function func(s *Subcommand)
}

func NewSubCommand(name string, fn func(s *Subcommand), function func(s *Subcommand)) *Subcommand {
	s := &Subcommand{
		Fs:       flag.NewFlagSet(name, flag.ContinueOnError),
		Name:     name,
		Function: function,
	}
	fn(s)
	return s
}

func (s *Subcommand) Init(args []string) error {
	return s.Fs.Parse(args)
}

func (s *Subcommand) Run() error {
	s.Function(s)
	return nil
}
