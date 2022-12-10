package tester

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/shivuslr41/grpc-tester/jq"
)

type Card struct {
	ID          string        `json:"id"`
	Description string        `json:"description"`
	Request     []interface{} `json:"requests"`
	Response    []interface{} `json:"responses"`
}

func (t *T) print(req []interface{}) {
	str := jq.Format(string(t.Response))
	var istr []interface{}
	err := json.Unmarshal([]byte(str), &istr)
	if err != nil {
		panic(err)
	}
	card := Card{
		ID:          t.ID,
		Description: t.Description,
		Request:     req,
		Response:    istr,
	}
	b, err := json.MarshalIndent(card, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func (t *T) test(e *Endpoint) error {
	if t.Skip {
		return nil
	}
	e.testerCall = true
	e.Data = t.Request
	e.StreamPayload = t.StreamPayload
	e.GrpcurlFlags = t.GrpcurlFlags
	reader, err := e.Run()
	if err != nil {
		return err
	}
	t.Response, err = io.ReadAll(reader)
	if err != nil {
		return err
	}
	if t.Compare {
		// compare logic here
	} else {
		t.print(e.Data)
	}
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

func RunTests(endpoints []Endpoint) {
	for i := range endpoints {
		if err := endpoints[i].Test(); err != nil {
			log.Fatal(err)
		}
	}
}
