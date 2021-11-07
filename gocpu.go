package main

import (
	"Github/sarthakpranesh/gocpu/utils"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

func main() {
	var cmds []*utils.Subcommand = []*utils.Subcommand{
		utils.NewSubCommand(
			"watch",
			func(s *utils.Subcommand) {
				s.Fs.IntVar(&s.Interval, "int", 2, "Usage --int: takes integer value for updating cpu freqency value")
			},
			WatchFrequency,
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
	log.Fatalln(string(usage))
}

func WatchFrequency(s *utils.Subcommand) {
	interval := s.Interval
	fileInfos, err := ioutil.ReadDir("/sys/devices/system/cpu/")
	if err != nil {
		log.Fatalln(err)
	}
	reg, err := regexp.Compile("cpu[0-9]{1,}")
	if err != nil {
		log.Fatalln(err)
	}
	var cpuCores []string
	for _, fInfo := range fileInfos {
		var name string = fInfo.Name()
		if reg.MatchString(name) {
			cpuCores = append(cpuCores, name)
		}
	}
	var cpuFreq []string = make([]string, len(cpuCores))
	var wg sync.WaitGroup
	wg.Add(1)
	utils.SetupCloseHandler()
	go func() {
		for {
			for i, core := range cpuCores {
				var fFreq string = "/sys/devices/system/cpu/" + core + "/cpufreq/scaling_cur_freq"
				freq, err := ioutil.ReadFile(fFreq)
				if err != nil {
					freq = []byte("0")
				}
				cpuFreq[i] = strings.ReplaceAll(string(freq), "\n", "")
			}
			wg.Done()
			time.Sleep(time.Second * time.Duration(interval))
		}
	}()
	for {
		wg.Wait()
		utils.ClearTerm()
		fmt.Printf("Total CPUs: %v\t(%v seconds)\n", len(cpuCores), interval)
		var fStringCpu string
		for i, cpuCore := range cpuCores {
			var cpu string = cpuCore + ":  " + cpuFreq[i]
			// normal
			fStringCpu = fStringCpu + cpu + "\t"
			if i%3 == 2 {
				fStringCpu = fStringCpu + "\n"
			}
		}
		fmt.Println(fStringCpu)
		wg.Add(1)
	}
}
