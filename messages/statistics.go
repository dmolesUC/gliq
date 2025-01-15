package messages

import "github.com/dmolesUC/gliq/params"

type Counts struct {
	All    int64
	Closed int64
	Opened int64
}

func (counts Counts) Included() int64 {
	var count int64
	switch params.StatesToInclude() {
	case params.Closed:
		count = counts.Closed
	case params.Opened:
		count = counts.Opened
	case params.All:
		count = counts.All
	}
	return count
}

type Statistics struct {
	Counts Counts
}
