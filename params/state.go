package params

import (
	"github.com/dmolesUC/gliq/config"
	"github.com/dmolesUC/gliq/util"
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
	var s State
	if config.IncludeOpen {
		s = s | Opened
	}
	if config.IncludeClosed {
		s = s | Closed
	}
	if s == 0 {
		util.Fail("can't return issues that are neither open nor closed")
	}
	return s
}
