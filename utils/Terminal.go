package utils

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

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

func CmdTerm(c string) {
	cmd := exec.Command("bash", "-c", c)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
