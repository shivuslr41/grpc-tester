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
	compare       bool
	print         bool
	global        bool
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
		flag.StringVarP(&data, "data", "d", "", "request in json format - '{\"name\":\"Ramesh\"}'")
		flag.BoolVarP(&streamPayload, "stream-payload", "S", false, "send multiple messages to server")
		flag.StringVarP(&grpcurlFlags, "grpcurl-flags", "g", "", "pass additional grpcurl flags - '-H \"Authorization: <TOKEN>\"'")
		flag.BoolVarP(&compare, "compare", "c", false, "test/compare responses")
		flag.BoolVarP(&print, "print", "P", false, "prints result")
		flag.BoolVarP(&global, "global", "G", false, "consider global flags for run/test commands")
	}
	flag.Parse()
}

func usage() {
	details := `
FORMAT:
	./grpc-tester [COMMAND] [FLAGS]

EXAMPLE:
	./grpc-tester list --server mygrpcserver:443 --tls

COMMANDS:
	gen		generates sample json.
	list		list services and methods.
	run		run executes requests provided in json or via -d flag.
	test		test responses againt expectations set.
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
		lister.Execute()
		return
	case "run", "test":
		if jsonFile != "" {
			command = "test"
		} else if data == "" {
			fmt.Println("--data | -d data is given empty!")
			fmt.Println("              OR")
			fmt.Println("--json | -j json file is not provided!")
			usage()
		}
	default:
		fmt.Println("given empty/invalid command!")
		usage()
	}

	// config global flags
	if global {
		tester.GConf.Use = global
		tester.GConf.Lister = *lister
		tester.GConf.Endpoint = endpoint
		tester.GConf.StreamPayload = streamPayload
		tester.GConf.GrpcurlFlags = grpcurlFlags
		tester.GConf.Compare = compare
		tester.GConf.Print = print
	}

	switch command {
	case "run":
		// construct runner
		runner := &tester.Runner{
			Lister:        *lister,
			Endpoint:      endpoint,
			StreamPayload: streamPayload,
			GrpcurlFlags:  grpcurlFlags,
			Data:          append([]interface{}{}, data),
		}
		validateCommandOptions(runner)
		runner.Execute()
		return
	case "test":
		endpoints := readJSON(jsonFile)
		validateCommandOptions(endpoints)
		tester.Execute(endpoints)
	}
}
