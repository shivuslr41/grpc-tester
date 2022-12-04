package tester

import (
	"fmt"
	"log"

	"github.com/shivuslr41/grpc-tester/exec"
)

func (l *Lister) Print() {
	sm, err := l.List()
	if err != nil {
		log.Fatal(err)
	}
	for s, m := range sm {
		fmt.Println("----------------<", s, ">-----------------")
		for i := range m {
			fmt.Println("----->", m[i])
		}
		fmt.Println()
	}
}

func (l *Lister) tlsFlag() string {
	var tls string
	if !l.TLS {
		tls = "--plaintext"
	}
	return tls
}

func (l *Lister) protoFlag() string {
	var protoOption string
	if l.ProtoPath != "" {
		protoOption = fmt.Sprintf("--import-path %s --proto %s", l.ProtoPath, l.ProtoFile)
	}
	return protoOption
}

func (l *Lister) List() (map[string][]string, error) {
	var server string
	if l.ProtoPath != "" {
		server = l.protoFlag()
	} else {
		server = l.Server
	}
	exe := exec.NewExec()
	exe.Flags = fmt.Sprintf("%s %s list", l.tlsFlag(), server)
	services, err := exe.GetCombinedStdout()
	if err != nil {
		return nil, err
	}
	servicesAndMethods := make(map[string][]string)
	for _, service := range services {
		exe.Flags = fmt.Sprintf("%s %s list %s", l.tlsFlag(), server, service)
		methods, err := exe.GetCombinedStdout()
		if err != nil {
			return nil, err
		}
		servicesAndMethods[service] = methods
	}
	return servicesAndMethods, nil
}
