package util

import (
	"fmt"
	"os"
)

func Fail(s string) {
	err := fmt.Errorf(s)
	_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(1)
}
