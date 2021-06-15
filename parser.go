package main

import (
	"bufio"
	"io"
	"strings"
)

const VALID_TOKENS = "><+-.,[]"

func Tokenize(handle io.Reader) (ret []string) {
	scanner := bufio.NewScanner(handle)

	for scanner.Scan() {
		sliced := strings.Split(scanner.Text(), "")

		for _, char := range sliced {
			if strings.Contains(VALID_TOKENS, char) {
				ret = append(ret, char)
			}
		}
	}

	return
}
