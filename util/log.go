package util

import (
	"fmt"
	"os"
)

func Log(a ...any) {
	// TODO: just use log.Printf
	_, _ = fmt.Fprintln(os.Stderr, a...)
}
