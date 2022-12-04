package main

import (
	"fmt"
	"os"

	tester "github.com/shivuslr41/grpc-tester"
	flag "github.com/spf13/pflag"
)

var (
	serverAddr string
	endpoint   string
	protoPath  string
	protoFile  string
	enableTLS  bool
	jsonFile   string
	help       bool
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
	case "list":
		validateLister(lister)
		lister.Print()
		return
	case "run", "test":
	default:
		fmt.Println("given empty/invalid command!")
		usage()
	}

	// TODO impl
	switch command {
	case "run":
		fmt.Println("running", serverAddr)
	case "test":
		fmt.Println("testing", serverAddr)
	}
}

// validate user options
func validateLister(l *tester.Lister) {
	if l.Server == "" && (l.ProtoPath == "" || l.ProtoFile == "") {
		fmt.Println("-s | --server is empty!")
		fmt.Println("-p | --proto-path is empty / -f | --proto-file is empty")
		usage()
	}
}
