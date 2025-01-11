package params

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/config"
)

type State uint8

const (
	Opened State = 1 << iota
	Closed
	All = Opened | Closed
)

var stateValues = map[State]string{
	Opened: "opened",
	Closed: "closed",
	All:    "all",
}

// StateVal parameter value for issue State (opened/closed/all)
func StateVal() string {
	s := StatesToInclude()
	return stateValues[s]
}

func StatesToInclude() State {
	s, err := configuredState()
	cobra.CheckErr(err)

	return s
}

func configuredState() (State, error) {
	var s State
	var err error
	if config.IncludeOpen {
		s = s | Opened
	}
	if config.IncludeClosed {
		s = s | Closed
	}
	if s == 0 {
		err = fmt.Errorf("can't return issues that are neither open nor closed")
	}
	return s, err
}
