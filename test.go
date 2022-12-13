package tester

import (
	"fmt"
	"io"
)

func (t *T) test(e *Endpoint) error {
	if t.Skip {
		return nil
	}

	e.testerCall = true
	e.Data = t.Request
	e.StreamPayload = t.StreamPayload
	e.GrpcurlFlags = t.GrpcurlFlags

	err := e.Run(func(rc io.ReadCloser) error {
		b, err := io.ReadAll(rc)
		if err != nil {
			return err
		}
		t.Response = b
		return nil
	})
	if err != nil {
		return err
	}

	if t.Print {
		t.print(e.Data)
	}

	if t.Compare {
		err = t.compare()
		if err != nil {
			return err
		}
	}

	fmt.Println("                                   -----------------------------------")
	return nil
}

func (e *Endpoint) Test() error {
	for i := range e.Tests {
		err := e.Tests[i].test(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func Execute(endpoints []Endpoint) {
	for i := range endpoints {
		if err := endpoints[i].Test(); err != nil {
			printErrAndExit(err)
		}
		fmt.Println("                                   ===================================")
	}
}
