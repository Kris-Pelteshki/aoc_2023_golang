package util

import "strings"

func SplitLines(input string) []string {
	return strings.Split(input, "\n")
}
