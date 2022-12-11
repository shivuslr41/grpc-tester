package tester

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/shivuslr41/grpc-tester/jq"
)

type Card struct {
	ID          string        `json:"id"`
	Description string        `json:"description"`
	Request     []interface{} `json:"requests"`
	Response    []interface{} `json:"responses"`
}

func (l *Lister) print() {
	sm, err := l.List()
	if err != nil {
		printErrAndExit(err)
	}
	for s, m := range sm {
		fmt.Println("         <<<", s, ">>>")
		for i := range m {
			fmt.Println("----->", m[i])
		}
		fmt.Println()
	}
}

func (r *Runner) print() {
	err := r.Run(func(rc io.ReadCloser) error {
		var out interface{}
		decoder := json.NewDecoder(rc)
		for decoder.More() {
			err := decoder.Decode(&out)
			if err != nil {
				return err
			}
			b, err := json.MarshalIndent(out, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		}
		return nil
	})
	if err != nil {
		printErrAndExit(err)
	}
}

func (t *T) print(req []interface{}) {
	str, err := jq.Format(string(t.Response))
	if err != nil {
		printErrAndExit(err)
	}
	var istr []interface{}
	err = json.Unmarshal([]byte(str), &istr)
	if err != nil {
		printErrAndExit(err)
	}
	b, err := json.MarshalIndent(
		Card{
			ID:          t.ID,
			Description: t.Description,
			Request:     req,
			Response:    istr,
		},
		"",
		"  ",
	)
	if err != nil {
		printErrAndExit(err)
	}
	fmt.Println(string(b))
}
