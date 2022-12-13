package tester

import (
	"encoding/json"
	"os"
)

func Generate() {
	var sampleJSON []Endpoint
	sampleJSON = append(sampleJSON, Endpoint{})
	sampleJSON[0].Tests = append(sampleJSON[0].Tests, T{})
	var dummyiFace []interface{}
	sampleJSON[0].Tests[0].Request = append(sampleJSON[0].Tests[0].Request, dummyiFace)
	sampleJSON[0].Tests[0].Expectations = append(sampleJSON[0].Tests[0].Expectations, dummyiFace)
	b, err := json.MarshalIndent(sampleJSON, "", "    ")
	if err != nil {
		printErrAndExit(err)
	}
	err = os.WriteFile("format.json", b, 0644)
	if err != nil {
		printErrAndExit(err)
	}
}
