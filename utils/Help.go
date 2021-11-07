package utils

import "fmt"

func Help(cmds []*Subcommand) {
	fmt.Printf("A simple cli tool to handle and watch your CPU \n\n")
	fmt.Println("Usage gocpu [subcommand] [flags]")
	fmt.Println("Commands ")
	fmt.Println("\twatch - see the realtime cpu frequency, updated at 2 seconds by default, can be changed using the -int flag")
	fmt.Println("\tturbo - sets the turbo on/off depending on the -enable flag")
	fmt.Println("\tgovern - lets you interactively select the cpu governor out the all the available governors")
	fmt.Printf("\nFlags \n")
	fmt.Println("\t-int - used with watch subcommand, value is treated a number of seconds between refreshes")
	fmt.Println("\t-enable - used with turbo subcommand, if passed turbo mode will be enabled else turbo mode will be disabled")
}
