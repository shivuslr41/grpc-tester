package tester

import (
	"errors"
	"io"
)

const (
	ErrTestsFailed = " ❌some tests failed!❌ "
)

// Test calls run to collect grpc result and compare them with expectations if set
func (t *T) Test(r Runner) error {
	// skips the test if set
	if t.Skip {
		return nil
	}

	// if global -G flag is set then override test file config
	if GConf.Use {
		if GConf.GrpcurlFlags != "" {
			t.GrpcurlFlags = GConf.GrpcurlFlags
		}
		t.Compare = GConf.Compare
		t.Print = GConf.Print
	}

	// create grpc requests from extracted data
	var err error
	if r.Data, err = t.replace(t.Requests); err != nil {
		return err
	}
	r.testerCall = !r.testerCall
	r.GrpcurlFlags = t.GrpcurlFlags

	// start grpc requests and collect responses
	// extract data from response if enabled
	// note: all streaming results are collects as combined/one response and compared further
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

	// compare the responses using jq
	if t.Compare {
		t.compare()
	}

	// print the test outcomes into console
	t.print()
	return nil
}

// test executes all test cases of an endpoint
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

// Execute starts test command
func Execute(endpoints []Endpoint) {
	// load previous results from json file into variables map
	if err := load(); err != nil {
		printErrAndExit(err)
	}
	for i := range endpoints {
		if err := endpoints[i].test(); err != nil {
			printErrAndExit(err)
		}
	}
	// save current results from variables map into json file.
	if err := save(); err != nil {
		printErrAndExit(err)
	}

	// check if overall test is passed or failed,
	// if any one test failed then exit with code 1
	// hence making overall tests failed.
	if overallFail {
		printErrAndExit(errors.New(ErrTestsFailed))
	}
}
