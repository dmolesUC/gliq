package options

import "github.com/dmolesUC/gliq/util"

// ------------------------------------------------------------
// Exported

type State uint8

const (
	Opened State = 1 << iota
	Closed
	All = Opened | Closed
)

func (s State) ToParam() string {
	return stateParams[s]
}

// ------------------------------------------------------------
// Package-local

var stateParams = map[State]string{
	Opened: "opened",
	Closed: "closed",
	All:    "all",
}

func toState(includeOpen bool, includeClosed bool) State {
	var s State
	if includeOpen {
		s = s | Opened
	}
	if includeClosed {
		s = s | Closed
	}
	if s == 0 {
		util.Fail("can't return issues that are neither open nor closed")
	}
	return s
}
