package util

import (
	"io"

	"github.com/spf13/cobra"
)

func QuietlyClose(c io.Closer) {
	err := c.Close()
	cobra.CheckErr(err)
}
