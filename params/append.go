package params

import (
	"net/url"

	"github.com/dmolesUC/gliq/config"
)

func ToRawQuery() string {
	query := url.Values{}
	query.Set("state", State())

	// with an authenticated user, GitLab by default returns only
	// issues created by that user
	if token := config.Token; len(token) > 0 {
		query.Set("scope", "all")
	}

	queryString := query.Encode()

	// GitLab wants raw commas, not %2F
	labels := Labels()
	if len(labels) > 0 {
		queryString = queryString + "&labels=" + Labels()
	}
	return queryString
}
