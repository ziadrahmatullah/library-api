package util

import "strings"

func RemoveNewLine(str string) string {
	return strings.Trim(str, "\n")
}