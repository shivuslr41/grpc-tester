package tester

import (
	"encoding/json"
	"fmt"

	"github.com/shivuslr41/grpc-tester/jq"
)

func (t *T) compare() error {
	result, err := jq.Format(string(t.Response))
	if err != nil {
		return err
	}

	b, err := json.Marshal(t.Expectations)
	if err != nil {
		return err
	}
	expect := string(b)

	filteredResult := result
	if len(t.Queries) == 0 {
		t.Queries = append(t.Queries, "'.'")
	}
	for i := range t.Queries {
		filteredResult, err = jq.Filter(filteredResult, t.Queries[i])
		if err != nil {
			return err
		}
	}

	// TODO: print if debug enabled.
	// fmt.Println(filteredResult)
	// fmt.Println(expect)

	pass, err := jq.Compare(filteredResult, expect, t.IgnoreOrder)
	if err != nil {
		return err
	}

	if pass {
		fmt.Println("PASS |", t.ID, "|", t.Description)
	} else {
		fmt.Println("FAIL |", t.ID, "|", t.Description)
	}
	return nil
}
