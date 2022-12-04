package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tester "github.com/shivuslr41/grpc-tester"
)

func readJSON(filename string) []tester.Endpoint {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var endpoints []tester.Endpoint
	if err = json.Unmarshal(b, &endpoints); err != nil {
		log.Fatal(err)
	}
	return endpoints
}

// validate user options
func validateCommandOptions(o interface{}) {
	switch op := o.(type) {
	case *tester.Lister:
		if isAnyOfEmpty(op.Server, op.ProtoPath, op.ProtoFile) {
			usage()
		}
	case *tester.Runner:
		if len(op.Data) == 0 {
			if isAnyOfEmpty(op.Server, op.ProtoPath, op.ProtoFile) {
				usage()
			}
		} else if op.Endpoint == "" {
			fmt.Println("--endpoint | -e is not provided!")
			usage()
		}
	case []tester.Endpoint:
		if len(op) == 0 || len(op[0].Tests) == 0 {
			fmt.Println("invalid json, no tests/requests provided!")
			os.Exit(0)
		}
	default:
		log.Fatal("invalid command type")
	}
}

func isAnyOfEmpty(svr, pp, pf string) bool {
	if svr == "" && (pp == "" || pf == "") {
		fmt.Println("-s | --server is empty!")
		fmt.Println("-p | --proto-path is empty / -f | --proto-file is empty")
		return true
	}
	return false
}
