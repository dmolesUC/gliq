package params

import (
	"strings"

	"github.com/dmolesUC/gliq/config"
)

func Labels() string {
	return strings.Join(config.Labels, ",")
}
