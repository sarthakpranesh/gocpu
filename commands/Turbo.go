package commands

import (
	"fmt"

	"github.com/sarthakpranesh/gocpu/utils"
)

func TurboSet(s *utils.Subcommand) {
	fmt.Println("Requesting sudo privilege")
	if s.Turbo {
		utils.CmdTerm("echo 0 | sudo tee /sys/devices/system/cpu/intel_pstate/no_turbo")
		fmt.Println("Enabling Turbo")
	} else {
		utils.CmdTerm("echo 1 | sudo tee /sys/devices/system/cpu/intel_pstate/no_turbo")
		fmt.Println("Disableing Turbo")
	}
}
