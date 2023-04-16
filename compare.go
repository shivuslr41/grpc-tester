package tester

import (
	"encoding/json"

	"github.com/shivuslr41/grpc-tester/jq"
)

// compare the response from a request with the expected output using the jq package.
// It filters the response data based on one or more queries provided by the user,
// then compares the filtered data with the expected data using the Compare function from the jq package,
// and sets the Pass flag accordingly.
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

	// prints if debug enabled
	print("Filtered result: " + filteredResult)
	print("Expected result: " + expect)

	t.Pass, err = jq.Compare(filteredResult, expect, t.IgnoreOrder)
	return err
}
