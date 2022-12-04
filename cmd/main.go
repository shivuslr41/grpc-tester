package main

import (
	"fmt"
	"os"

	tester "github.com/shivuslr41/grpc-tester"
	flag "github.com/spf13/pflag"
)

var (
	serverAddr    string
	endpoint      string
	protoPath     string
	protoFile     string
	enableTLS     bool
	jsonFile      string
	help          bool
	data          string
	streamPayload bool
	grpcurlFlags  string
)

func init() {
	{
		flag.StringVarP(&serverAddr, "server", "s", "", "gRPC server address")
		flag.StringVarP(&endpoint, "endpoint", "e", "", "service and method to call")
		flag.StringVarP(&protoPath, "proto-path", "p", "", "proto path, if server doesn't support grpc reflection")
		flag.StringVarP(&protoFile, "proto-file", "f", "", "proto file")
		flag.BoolVarP(&enableTLS, "tls", "t", false, "use tls connection")
		flag.StringVarP(&jsonFile, "json", "j", "", "json file containing test scopes")
		flag.BoolVarP(&help, "help", "h", false, "shows tool usage")
		flag.StringVarP(&data, "data", "d", "", "request in json format - '{\"name\":\"Bob\"}'")
		flag.BoolVarP(&streamPayload, "stream-payload", "m", false, "send multiple messages to server")
		flag.StringVarP(&data, "grpcurl-flags", "g", "", "pass additional grpcurl flags - '-H Authorization: <TOKEN>'")
	}
	flag.Parse()
}

func usage() {
	details := `
FORMAT:
	./tester [COMMAND] [FLAGS]

EXAMPLE:
	./tester list --server mygrpcserver:443 --tls

COMMANDS:
	gen		generates sample json.
	list		list services and methods.
	run		run requests provided in json.
	test		test responses againt expectation set.
`
	fmt.Println(details)
	flag.Usage()
	os.Exit(0)
}

func main() {
	// print usage
	if help {
		usage()
	}

	// get command to exec
	command := flag.Arg(0)

	// construct lister
	lister := &tester.Lister{
		Server:    serverAddr,
		ProtoPath: protoPath,
		ProtoFile: protoFile,
		TLS:       enableTLS,
	}

	// check command type
	switch command {
	case "gen":
		tester.Generate()
		return
	case "list":
		validateCommandOptions(lister)
		lister.Print()
		return
	case "run", "test":
		if jsonFile != "" {
			command = "test"
		} else if data == "" {
			fmt.Println("--data | -d data is given empty!")
			fmt.Println("--json | -j json file is not provided!")
			usage()
			return
		}
	default:
		fmt.Println("given empty/invalid command!")
		usage()
		return
	}

	// construct runner and tester
	runner := &tester.Runner{
		Lister:        *lister,
		Endpoint:      endpoint,
		StreamPayload: streamPayload,
		GrpcurlFlags:  grpcurlFlags,
	}
	if data != "" {
		runner.Data = append(runner.Data, data)
	}

	switch command {
	case "run":
		validateCommandOptions(runner)
		runner.Print()
		return
	case "test":
		endpoints := readJSON(jsonFile)
		validateCommandOptions(endpoints)
		tester.RunTests(endpoints)
	}
}
