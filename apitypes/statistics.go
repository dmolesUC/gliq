package apitypes

import "github.com/dmolesUC/gliq/options"

// ------------------------------------------------------------
// Counts

type Counts struct {
	All    int64
	Closed int64
	Opened int64
}

func (counts Counts) Included() int64 {
	var count int64
	switch options.StateFlags() {
	case options.Closed:
		count = counts.Closed
	case options.Opened:
		count = counts.Opened
	case options.All:
		count = counts.All
	}
	return count
}

// ------------------------------------------------------------
// Statistics

type Statistics struct {
	Counts Counts
}

// ------------------------------------------------------------
// Statistics response wrapper

type StatisticsResponse struct {
	Statistics Statistics
}
