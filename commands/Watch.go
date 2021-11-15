package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sarthakpranesh/gocpu/utils"
)

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
	var noTurboState int64
	var wg sync.WaitGroup
	wg.Add(2)
	utils.SetupCloseHandler()
	// For cpu frequency
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
	// For turbo enable/disable
	go func() {
		var noTurboFile string = "/sys/devices/system/cpu/intel_pstate/no_turbo"
		for {
			noTurboStateByte, err := ioutil.ReadFile(noTurboFile)
			if err != nil {
				noTurboStateByte = []byte("2")
			}
			noTurboStr := strings.ReplaceAll(string(noTurboStateByte), "\n", "")
			noTurboState, err = strconv.ParseInt(noTurboStr, 10, 64)
			if err != nil {
				noTurboState = 2
			}
			wg.Done()
			time.Sleep(time.Second * time.Duration(interval))
		}
	}()
	for {
		wg.Wait()
		utils.ClearTerm()
		fmt.Printf("Total CPUs: %v\t(%v seconds)\n", len(cpuCores), interval)
		fmt.Printf("Turbo Enabled: %v\t(%v no turbo)\n", noTurboState == 0, noTurboState) // turbo enabled if noTurbo is set to 0
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
		wg.Add(2)
	}
}
