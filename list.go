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

func (l *Lister) List() (map[string][]string, error) {
	var tls string
	if !l.TLS {
		tls = "--plaintext"
	}
	var server string
	if l.ProtoPath != "" {
		server = fmt.Sprintf("--import-path %s --proto %s", l.ProtoPath, l.ProtoFile)
	} else {
		server = l.Server
	}
	exe := exec.NewExec()
	exe.Flags = fmt.Sprintf("%s %s list", tls, server)
	services, err := exe.GetCombinedStdout()
	if err != nil {
		return nil, err
	}
	servicesAndMethods := make(map[string][]string)
	for _, service := range services {
		exe.Flags = fmt.Sprintf("%s %s list %s", tls, server, service)
		methods, err := exe.GetCombinedStdout()
		if err != nil {
			return nil, err
		}
		servicesAndMethods[service] = methods
	}
	return servicesAndMethods, nil
}
