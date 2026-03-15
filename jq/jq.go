package jq

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shivuslr41/grpc-tester/exec"
)

// shellSingleQuote returns a shell-safe single-quoted representation of s.
// It uses the POSIX pattern: 'foo'\''bar' to embed single quotes.
func shellSingleQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}

// Format a JSON string using the "jq -s" command-line tool.
func Format(j string) (string, error) {
	b, err := exec.NewCMD(
		fmt.Sprintf(
			"echo %s | jq -s '.'",
			shellSingleQuote(j),
		),
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s", string(b))
	}
	return string(b), nil
}

// Filter (select sub-objects) a JSON string using the "jq -c" command-line tool.
func Filter(j string, q string) (string, error) {
	b, err := exec.NewCMD(
		fmt.Sprintf(
			"echo %s | jq -c %s",
			shellSingleQuote(j),
			shellSingleQuote(q),
		),
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s", string(b))
	}
	return string(b), nil
}

// Compare two JSON strings using the "jq --argjson" command-line tool.
func Compare(r, e string, o bool) (bool, error) {
	var command string
	// If the "o" flag is set to true, the function sorts any arrays within the JSON strings before comparing them
	if o {
		command = fmt.Sprintf(
			"jq --argjson a %s --argjson b %s -n '($a | (.. | arrays) |= sort) as $a | ($b | (.. | arrays) |= sort) as $b | $a == $b'",
			shellSingleQuote(r),
			shellSingleQuote(e),
		)
	} else {
		command = fmt.Sprintf(
			"jq --argjson a %s --argjson b %s -n '($a == $b)'",
			shellSingleQuote(r),
			shellSingleQuote(e),
		)
	}
	b, err := exec.NewCMD(command).CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("%s", string(b))
	}
	return strconv.ParseBool(strings.Split(string(b), "\n")[0])
}

// Replace a portion of a JSON string using the "jq" command-line tool.
func Replace(o, q, d string) (string, error) {
	b, err := exec.NewCMD(
		fmt.Sprintf(
			"echo %s | jq -rc %s",
			shellSingleQuote(o),
			shellSingleQuote(fmt.Sprintf("%s |= %s", q, d)),
		),
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s", string(b))
	}
	return string(b), nil
}

// Extract (select sub-objects) a JSON string using the "jq -rc" command-line tool.
// similar to filter function but this returns raw format of JSON string.
func Extract(o, q string) (string, error) {
	b, err := exec.NewCMD(
		fmt.Sprintf(
			"echo %s | jq -rc %s",
			shellSingleQuote(o),
			shellSingleQuote(q),
		),
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s", string(b))
	}
	return string(b), nil
}
