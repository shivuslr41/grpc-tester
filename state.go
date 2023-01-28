// maintain variables state for multiple runs.
package tester

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shivuslr41/grpc-tester/jq"
)

func (s *State) replace(req []any) ([]any, error) {
	if len(s.Replace) == 0 || len(s.ReplaceFrom) == 0 {
		return req, nil
	}
	r, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	request := string(r)

	for i := range s.ReplaceFrom {
		if v, ok := variables[s.ReplaceFrom[i]]; ok {
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
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
	return req, nil
}

func (s *State) extract(res []any) error {
	if len(s.Extract) == 0 || len(s.ExtractTo) == 0 {
		return nil
	}
	response, err := json.Marshal(res)
	if err != nil {
		return err
	}
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
	}
	return nil
}
