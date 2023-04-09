package tester

import (
	"encoding/json"
	"os"
)

// The Generate function creates a sample JSON file with a pre-defined structure and writes it to a file named format.json.
// The JSON file can be used as a starting point for creating test cases.
func Generate() {
	var sampleJSON []Endpoint
	sampleJSON = append(sampleJSON, Endpoint{})
	sampleJSON[0].Tests = append(sampleJSON[0].Tests, T{})
	sampleJSON[0].Tests[0].Requests = append(sampleJSON[0].Tests[0].Requests, struct{}{})
	sampleJSON[0].Tests[0].Expectations = append(sampleJSON[0].Tests[0].Expectations, struct{}{})
	sampleJSON[0].Tests[0].Queries = append(sampleJSON[0].Tests[0].Queries, "")
	b, err := json.MarshalIndent(sampleJSON, "", "    ")
	if err != nil {
		printErrAndExit(err)
	}
	err = os.WriteFile("format.json", b, 0644)
	if err != nil {
		printErrAndExit(err)
	}
}
