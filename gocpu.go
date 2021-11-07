package main

import (
	"os"

	"github.com/sarthakpranesh/gocpu/commands"
	"github.com/sarthakpranesh/gocpu/utils"
)

func main() {
	var cmds []*utils.Subcommand = []*utils.Subcommand{
		utils.NewSubCommand(
			"watch",
			func(s *utils.Subcommand) {
				s.Fs.IntVar(&s.Interval, "int", 2, "Takes integer value for updating cpu freqency value")
			},
			commands.WatchFrequency,
		),
		utils.NewSubCommand(
			"turbo",
			func(s *utils.Subcommand) {
				s.Fs.BoolVar(&s.Turbo, "enable", false, "If not passed, turbo boost will be disabled")
			},
			commands.TurboSet,
		),
		utils.NewSubCommand(
			"govern",
			func(s *utils.Subcommand) {},
			commands.Governor,
		),
	}

	if len(os.Args) > 1 {
		subcommand := os.Args[1]
		for _, cmd := range cmds {
			if cmd.Name == subcommand {
				cmd.Init(os.Args[2:])
				cmd.Run()
				return
			}
		}
	}

	utils.Help(cmds)
}
