package util

import (
	"fmt"
	"os"
)

func Log(a ...any) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
}
