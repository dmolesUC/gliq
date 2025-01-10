package params

import (
	"fmt"

	"github.com/spf13/cobra"

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
func State() string {
	s, err := configuredState()
	cobra.CheckErr(err)

	return stateValues[s]
}

func configuredState() (state, error) {
	var s state
	var err error
	if config.IncludeOpen {
		s = s | opened
	}
	if config.IncludeClosed {
		s = s | closed
	}
	if s == 0 {
		err = fmt.Errorf("can't return issues that are neither open nor closed")
	}
	return s, err
}
