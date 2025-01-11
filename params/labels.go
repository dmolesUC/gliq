package params

import (
	"strings"

	"github.com/dmolesUC/gliq/config"
)

func LabelsVal() string {
	return strings.Join(config.Labels, ",")
}
