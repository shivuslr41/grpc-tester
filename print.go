package tester

import (
	"encoding/json"
	"fmt"
	"io"
)

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
	if err := r.Run(func(rc io.ReadCloser) error {
		var out interface{}
		decoder := json.NewDecoder(rc)
		for decoder.More() {
			if err := decoder.Decode(&out); err != nil {
				return err
			}
			b, err := json.MarshalIndent(out, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		}
		return nil
	}); err != nil {
		printErrAndExit(err)
	}
}

func (t *T) print() {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		printErrAndExit(err)
	}
	fmt.Println(string(b))
	if !t.Skip && t.Compare {
		if t.Pass {
			fmt.Println("PASS |", t.ID, "|", t.Description)
		} else {
			fmt.Println("FAIL |", t.ID, "|", t.Description)
		}
	}
	fmt.Println("                                   -----------------------------------")
}
