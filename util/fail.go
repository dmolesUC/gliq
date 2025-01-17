package util

import (
	"fmt"
)

func Fail(s string) {
	err := fmt.Errorf(s)
	QuietlyHandle(err)
}
