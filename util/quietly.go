package util

import (
	"fmt"
	"io"
	"os"
)

func QuietlyClose(c io.Closer) {
	err := c.Close()
	QuietlyHandle(err)
}

func QuietlyHandle(err error) {
	if err == nil {
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(1)
}

func Safely[T any](f func() (T, error)) T {
	v, err := f()
	QuietlyHandle(err)
	return v
}
