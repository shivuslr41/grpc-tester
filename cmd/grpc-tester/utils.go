package main

import (
	"encoding/json"
	"fmt"
	"os"

	tester "github.com/shivuslr41/grpc-tester"
)

// read and unmarshal JSON data from a file into a slice of Endpoint objects.
func readJSON(filename string) []tester.Endpoint {
	b, err := os.ReadFile(filename)
	if err != nil {
		printErrAndExit(err)
	}
	var endpoints []tester.Endpoint
	if err = json.Unmarshal(b, &endpoints); err != nil {
		printErrAndExit(err)
	}
	return endpoints
}

// validate user options
func validateCommandOptions(o any) {
	switch op := o.(type) {
	case *tester.Lister:
		if op.Server == "" && (op.ProtoPath == "" || op.ProtoFile == "") {
			fmt.Println("-s | --server address is not provided!")
			fmt.Println("                       OR")
			fmt.Println("-p | --proto-path is empty / -f | --proto-file is empty")
			usage()
		}
	case *tester.Runner:
		if op.Server == "" {
			fmt.Println("-s | --server address is not provided!")
			usage()
		}
		if op.Endpoint == "" {
			fmt.Println("--endpoint | -e is not provided!")
			usage()
		}
	case []tester.Endpoint:
		if len(op) == 0 || len(op[0].Tests) == 0 {
			fmt.Print("invalid json, no tests/requests provided!")
			os.Exit(0)
		}
	default:
		fmt.Println("invalid command type")
		usage()
	}
}

// prints error and exit
func printErrAndExit(err error) {
	fmt.Print(err)
	os.Exit(1)
}
