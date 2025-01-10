package util

import "strings"

func Prettify(s string) string {
	return strings.Join(strings.Split(s, "\n\t"), "\n")
}
