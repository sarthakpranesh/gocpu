package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"
)

var cmdTypes []string = []string{"watch"}

func main() {
	// useColor := flag.Bool("color", false, "display colorized output")
	watchInterval := flag.Int("int", 2, "Usage --int: takes integer value for updating cpu freqency value")
	flag.Parse()
	var curCmd string
	for _, cmd := range cmdTypes {
		if cmd == flag.Arg(0) {
			curCmd = cmd
		}
	}

	// if no usage match
	if curCmd == "" {
		usage, _ := ioutil.ReadFile("usage.txt")
		log.Fatalln(string(usage))
		return
	}

	// if usage matched
	if curCmd == "watch" {
		WatchFrequency(*watchInterval)
	}
}

func WatchFrequency(interval int) {
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
	SetupCloseHandler()
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
			time.Sleep(time.Second)
		}
	}()
	for {
		wg.Wait()
		ClearTerm()
		fmt.Printf("Total CPUs: %v\n", len(cpuCores))
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

func ClearTerm() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\rCtrl+C pressed, exiting real time frequency watch")
		os.Exit(0)
	}()
}
