package jq

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shivuslr41/grpc-tester/exec"
)

// format a JSON string using the "jq -s" command-line tool.
func Format(j string) (string, error) {
	b, err := exec.NewCMD(
		fmt.Sprintf(
			"echo '%s' | jq -s '.'",
			j,
		),
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s", string(b))
	}
	return string(b), nil
}

// filter(select sub-objects) a JSON string using the "jq -c" command-line tool.
func Filter(j string, q string) (string, error) {
	b, err := exec.NewCMD(
		fmt.Sprintf(
			"echo '%s' | jq -c %s",
			j,
			q,
		),
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s", string(b))
	}
	return string(b), nil
}

// compare two JSON strings using the "jq --argjson" command-line tool.
func Compare(r, e string, o bool) (bool, error) {
	var command string
	// If the "o" flag is set to true, the function sorts any arrays within the JSON strings before comparing them
	if o {
		command = fmt.Sprintf(
			"jq --argjson a '%s' --argjson b '%s' -n '($a | (.. | arrays) |= sort) as $a | ($b | (.. | arrays) |= sort) as $b | $a == $b'",
			r,
			e,
		)
	} else {
		command = fmt.Sprintf(
			"jq --argjson a '%s' --argjson b '%s' -n '($a == $b)'",
			r,
			e,
		)
	}
	b, err := exec.NewCMD(command).CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("%s", string(b))
	}
	return strconv.ParseBool(strings.Split(string(b), "\n")[0])
}

// replace a portion of a JSON string using the "jq" command-line tool.
func Replace(o, q, d string) (string, error) {
	b, err := exec.NewCMD(
		fmt.Sprintf(
			"echo '%s' | jq -rc '%s |= %s'",
			o,
			q,
			d,
		),
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s", string(b))
	}
	return string(b), nil
}

// extract(select sub-objects) a JSON string using the "jq -rc" command-line tool.
// similar to filter function but this returns raw format of JSON string.
func Extract(o, q string) (string, error) {
	b, err := exec.NewCMD(
		fmt.Sprintf(
			"echo '%s' | jq -rc '%s'",
			o,
			q,
		),
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s", string(b))
	}
	return string(b), nil
}
