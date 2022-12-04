package tester

import (
	"encoding/json"
	"log"
	"os"
)

func Generate() {
	var sampleJSON Tester
	sampleJSON.Tests = append(sampleJSON.Tests, T{})
	b, err := json.MarshalIndent(sampleJSON, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("format.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
