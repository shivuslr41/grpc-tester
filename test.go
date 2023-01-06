package tester

import (
	"io"
)

func (t *T) Test(r Runner) error {
	if t.Skip {
		return nil
	}

	r.testerCall = true
	r.Data = t.Requests
	r.GrpcurlFlags = t.GrpcurlFlags

	err := r.Run(func(rc io.ReadCloser) error {
		b, err := io.ReadAll(rc)
		if err != nil {
			return err
		}
		// format response
		err = t.format(b)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	if t.Compare {
		return t.compare()
	}

	if t.Print {
		t.print()
	}
	return nil
}

func (e *Endpoint) test() error {
	for i := range e.Tests {
		if err := e.Tests[i].Test(e.Runner); err != nil {
			return err
		}
	}
	return nil
}

func Execute(endpoints []Endpoint) {
	for i := range endpoints {
		if err := endpoints[i].test(); err != nil {
			printErrAndExit(err)
		}
	}
}
