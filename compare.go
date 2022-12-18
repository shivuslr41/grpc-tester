package tester

import (
	"encoding/json"

	"github.com/shivuslr41/grpc-tester/jq"
)

func (t *T) compare() error {
	b, err := json.Marshal(t.Response)
	if err != nil {
		return err
	}
	filteredResult := string(b)

	b, err = json.Marshal(t.Expectations)
	if err != nil {
		return err
	}
	expect := string(b)

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

	t.Pass, err = jq.Compare(filteredResult, expect, t.IgnoreOrder)
	return err
}
