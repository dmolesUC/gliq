package params

import (
	"net/url"
	"strings"

	"github.com/dmolesUC/gliq/config"
)

func LabelsVal() string {
	return strings.Join(config.Labels, ",")
}

// TODO: move this to its own file
func MilestonesVal() string {
	return url.QueryEscape(config.Milestone)
}
