// maintain variables state for multiple runs.
package tester

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shivuslr41/grpc-tester/jq"
)

// replaces current test file grpc requests with previous responses from different test files/cases.
// particularly using when requests needs to formed dynamically depending on the previous grpc calls results.
// example: one grpc call returned result "id" needed to be used in request of another grpc call.
func (s *State) replace(req []any) ([]any, error) {
	// if replace capability is not set then return original request back
	if len(s.Replace) == 0 || len(s.ReplaceFrom) == 0 {
		return req, nil
	}
	r, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	request := string(r)

	// replaceFrom contains JSON field that needed to be extracted from stored variables file.
	for i := range s.ReplaceFrom {
		if v, ok := variables[s.ReplaceFrom[i]]; ok {
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			// replace extracted part into current request.
			request, err = jq.Replace(request, s.Replace[i], string(b))
			if err != nil {
				return nil, err
			}
		} else {
			fmt.Println("extracted data not found for", s.ReplaceFrom[i])
		}
	}
	if err = json.Unmarshal([]byte(request), &req); err != nil {
		return nil, err
	}
	// prints if debug enabled
	print("Request:", req)
	return req, nil
}

// extract part of/whole data from grpc result and stores in global variables map.
// similar to replace method, extract works on the result area.
// extracted data can be further used other grpc calls request.
func (s *State) extract(res []any) error {
	// ignore storing anything into variables map if disabled
	if len(s.Extract) == 0 || len(s.ExtractTo) == 0 {
		return nil
	}
	response, err := json.Marshal(res)
	if err != nil {
		return err
	}

	// extract the specified JSON fields from the response and store into variables map
	for i := range s.Extract {
		str, err := jq.Extract(string(response), s.Extract[i])
		if err != nil {
			return err
		}
		str = strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					str, "\t", "",
				), "\r", "",
			), "\n", "",
		)
		variables[s.ExtractTo[i]] = str
		// prints if debug enabled
		print("Extracted:", str)
	}
	return nil
}
