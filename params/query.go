package params

import (
	"net/url"

	"github.com/dmolesUC/gliq/config"
)

func ToRawQuery() string {
	query := url.Values{}
	setState(query)
	setScope(query)
	setMilestone(query)

	return setLabels(query.Encode())
}

func setState(query url.Values) {
	query.Set("State", StateVal())
}

func setScope(query url.Values) {
	// TODO: put this code and Authentication header code together

	// with an authenticated user, GitLab by default returns only
	// issues created by that user
	if token := config.Token; len(token) > 0 {
		query.Set("scope", "all")
	}
}

func setMilestone(query url.Values) {
	if milestone := MilestonesVal(); len(milestone) > 0 {
		query.Set("milestone", milestone)
	}
}

func setLabels(queryString string) string {
	// GitLab wants raw commas, not %2F
	labels := LabelsVal()
	if len(labels) > 0 {
		queryString = queryString + "&labels=" + LabelsVal()
	}
	return queryString
}
