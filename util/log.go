package util

import (
	"fmt"
	"os"
	"strings"
)

func Log(a ...any) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
}

func Prettify(s string) string {
	return strings.Join(strings.Split(s, "\n\t"), "\n")
}
