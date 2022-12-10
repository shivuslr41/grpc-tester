package tester

import (
	"encoding/json"
	"log"
	"os"
)

func Generate() {
	var sampleJSON []Endpoint
	sampleJSON = append(sampleJSON, Endpoint{})
	sampleJSON[0].Tests = append(sampleJSON[0].Tests, T{})
	var dummyFace []interface{}
	sampleJSON[0].Tests[0].Request = append(sampleJSON[0].Tests[0].Request, dummyFace)
	sampleJSON[0].Tests[0].Expectations = append(sampleJSON[0].Tests[0].Expectations, dummyFace)
	b, err := json.MarshalIndent(sampleJSON, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("format.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
