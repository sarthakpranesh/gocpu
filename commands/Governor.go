package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/sarthakpranesh/gocpu/utils"
)

func Governor(s *utils.Subcommand) {
	governorByte, err := ioutil.ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/scaling_available_governors")
	if err != nil {
		log.Fatalln("Can't find out available governors")
	}
	governorString := strings.ReplaceAll(string(governorByte), "\n", "")
	governors := strings.Split(governorString, " ")
	fmt.Println("Available governors:")
	for i, gov := range governors {
		fmt.Printf("%v)%v \t", i, gov)
	}
	fmt.Printf("\nEnter the Governor id to use: ")
	var ids string
	_, err = fmt.Scanln(&ids)
	if err != nil {
		log.Fatalln("Unable to take user input:", err)
	}
	id, err := strconv.ParseInt(ids, 10, 0)
	if err != nil {
		log.Fatalln("Invalid user input:", err)
	}
	var cmd string = "echo " + governors[id] + " | sudo tee /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor"
	utils.CmdTerm(cmd)
	fmt.Println("Governor set:", governors[id])
}
