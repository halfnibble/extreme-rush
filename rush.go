package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// RushConfig ...
type RushConfig struct {
	Version string `json:"rushVersion"`
}

func main() {
	// Validate schema
	rushSchema, err := parseSchema()
	rushFile, err := os.Open("test/rush.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = rushSchema.Validate(rushFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Parse test/rush.json
	var rushConfig RushConfig
	rushFileBytes, err := ioutil.ReadFile("test/rush.json")
	if err := json.Unmarshal(rushFileBytes, &rushConfig); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Rush version: " + rushConfig.Version)

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
