package tester

import (
	"fmt"
	"strings"

	"github.com/shivuslr41/grpc-tester/exec"
)

func (l *Lister) List() (map[string][]string, error) {
	var server string
	if l.ProtoPath != "" {
		server = l.protoFlag()
	} else {
		server = l.Server
	}
	b, err := exec.NewCMD(
		fmt.Sprintf("grpcurl %s %s list",
			l.tlsFlag(), server,
		),
	).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s", string(b))
	}
	servicesAndMethods := make(map[string][]string)
	for _, service := range removeEmptyStrings(strings.Split(string(b), "\n")) {
		b, err = exec.NewCMD(
			fmt.Sprintf("grpcurl %s %s list %s",
				l.tlsFlag(),
				server,
				service,
			),
		).CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("%s", string(b))
		}
		servicesAndMethods[service] = removeEmptyStrings(strings.Split(string(b), "\n"))
	}
	return servicesAndMethods, nil
}

func (l *Lister) Execute() {
	l.print()
}
