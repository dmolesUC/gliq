package params

import (
	"fmt"

	"github.com/dmolesUC/gliq/config"
)

type state uint8

const (
	opened state = 1 << iota
	closed
	all = opened | closed
)

var stateValues = map[state]string{
	opened: "opened",
	closed: "closed",
	all:    "all",
}

// State parameter value for issue state (opened/closed/all)
func State() (string, error) {
	var s state
	if config.IncludeOpen {
		s = s | opened
	}
	if config.IncludeClosed {
		s = s | closed
	}
	if str, ok := stateValues[s]; ok {
		return str, nil
	}
	return "", fmt.Errorf("can't return issues that are neither open nor closed")
}
