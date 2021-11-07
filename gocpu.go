package main

import (
	"Github/sarthakpranesh/gocpu/commands"
	"Github/sarthakpranesh/gocpu/utils"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var cmds []*utils.Subcommand = []*utils.Subcommand{
		utils.NewSubCommand(
			"watch",
			func(s *utils.Subcommand) {
				s.Fs.IntVar(&s.Interval, "int", 2, "Usage --int: takes integer value for updating cpu freqency value")
			},
			commands.WatchFrequency,
		),
		utils.NewSubCommand(
			"turbo",
			func(s *utils.Subcommand) {
				s.Fs.BoolVar(&s.Turbo, "enable", false, "Usage --enable: If not passed, turbo boost will be disabled")
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

	usage, _ := ioutil.ReadFile("usage.txt")
	fmt.Println(string(usage))
}
