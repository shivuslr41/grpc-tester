package jq

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shivuslr41/grpc-tester/exec"
)

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

func Compare(r, e string, o bool) (bool, error) {
	var command string
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
