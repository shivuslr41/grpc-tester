package tester

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func (t *T) print(req []interface{}) {
	result := `{
		"id":"` + t.ID + `",
		"description":"` + t.Description + `",
		"request":` + fmt.Sprint(req) + `,
		"response":` + string(t.Response) + `,
	}`
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("response:", string(t.Response))
		log.Fatal(err)
	}
	fmt.Println(string(b))
	// fmt.Println("----------------------------------------------------------")
	// fmt.Println("ID:", t.ID)
	// fmt.Println("Description", t.Description)
	// fmt.Printf("Request:\n%s\n", req)
	// fmt.Printf("Response:\n%s\n", t.Response)
	// fmt.Println("----------------------------------------------------------")
}

func (t *T) test(e *Endpoint) error {
	if t.Skip {
		return nil
	}
	e.Data = t.Request
	e.StreamPayload = t.StreamPayload
	reader, err := e.Run()
	if err != nil {
		return err
	}
	// some error here at readall - invalid character 'm' looking for beginning of value
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
