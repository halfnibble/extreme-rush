package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// list subcommand
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	testFlag := listCommand.Bool("test", false, "Does this work?")
	debugFlag := listCommand.Bool("debug", false, "Will print more info.")

	command := ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "list":
		listCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if listCommand.Parsed() {
		if *debugFlag {
			fmt.Println("Args: " + strings.Join(os.Args[1:], ", "))
		}

		if *testFlag {
			fmt.Println("It works!")
		} else {
			fmt.Println("No test flag set.")
		}
	}
}
