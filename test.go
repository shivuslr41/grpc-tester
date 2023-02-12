package tester

import (
	"io"
)

func (t *T) Test(r Runner) error {
	if t.Skip {
		return nil
	}

	if GConf.Use {
		if GConf.GrpcurlFlags != "" {
			t.GrpcurlFlags = GConf.GrpcurlFlags
		}
		t.Compare = GConf.Compare
		t.Print = GConf.Print
	}

	var err error
	if r.Data, err = t.replace(t.Requests); err != nil {
		return err
	}
	r.testerCall = !r.testerCall
	r.GrpcurlFlags = t.GrpcurlFlags

	err = r.Run(func(rc io.ReadCloser) error {
		b, err := io.ReadAll(rc)
		if err != nil {
			return err
		}
		// format response
		if err = t.format(b); err != nil {
			return err
		}
		// extract fields from response
		if err = t.extract(t.Response); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	if t.Compare {
		t.compare()
	}
	t.print()
	return nil
}

func (e *Endpoint) test() error {
	if e.Skip {
		return nil
	}
	for i := range e.Tests {
		if err := e.Tests[i].Test(e.Runner); err != nil {
			return err
		}
	}
	return nil
}

func Execute(endpoints []Endpoint) {
	if err := load(); err != nil {
		printErrAndExit(err)
	}
	for i := range endpoints {
		if err := endpoints[i].test(); err != nil {
			printErrAndExit(err)
		}
	}
	if err := save(); err != nil {
		printErrAndExit(err)
	}
}
