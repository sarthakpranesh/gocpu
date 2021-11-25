package utils_test

import (
	"testing"

	"github.com/sarthakpranesh/gocpu/utils"
)

func TestNewSubCommand(t *testing.T) {
	var fn func(s *utils.Subcommand) = func(s *utils.Subcommand) {
		s.Fs.BoolVar(&s.Turbo, "turbo", false, "Test flag not passed")
		s.Fs.IntVar(&s.Interval, "int", 2, "Integer not passed")
	}
	var function func(s *utils.Subcommand) = func(s *utils.Subcommand) {
		if s.Name != "test" {
			t.Error("Can't set subcommand name")
		}
		if s.Turbo != true {
			t.Error("Can't set `--turbo` flag with command")
		}
		if s.Interval != 1 {
			t.Error("Can't set '--int' flag with command")
		}
	}
	var subcommand *utils.Subcommand = utils.NewSubCommand("test", fn, function)
	var args []string = []string{"--turbo", "--int", "1"} // mocking os.Args
	subcommand.Init(args)
	subcommand.Run()
}
