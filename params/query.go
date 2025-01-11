package params

import (
	"net/url"

	"github.com/dmolesUC/gliq/config"
)

func ToRawQuery() string {
	query := url.Values{}
	query.Set("State", StateVal())

	// with an authenticated user, GitLab by default returns only
	// issues created by that user
	if token := config.Token; len(token) > 0 {
		query.Set("scope", "all")
	}

	if milestone := MilestonesVal(); len(milestone) > 0 {
		query.Set("milestone", milestone)
	}

	queryString := query.Encode()

	// GitLab wants raw commas, not %2F
	labels := LabelsVal()
	if len(labels) > 0 {
		queryString = queryString + "&labels=" + LabelsVal()
	}
	return queryString
}
