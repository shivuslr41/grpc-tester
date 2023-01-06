package tester

import (
	"encoding/json"
	"os"
)

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
