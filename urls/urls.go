package urls

import (
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/dmolesUC/gliq/options"
	"github.com/dmolesUC/gliq/util"
)

// ------------------------------------------------------------
// Exported

func RepoUrl() *ApiUrl {
	if repoUrl == nil {
		repo := url.PathEscape(options.Repository())
		repoUrl = apiBaseUrl.JoinPath("projects", repo)
	}
	return repoUrl
}

func IssuesUrl() *ApiUrl {
	return RepoUrl().JoinPath("issues")
}

func IssueStatsUrl() *ApiUrl {
	return RepoUrl().JoinPath("issues_statistics")
}

// ReadAs reads data from the specified API URL and marshals it to the
// specified type. (This is a top-level function rather than a method on
// ApiUrl because Go doesn't support parameterized methods, because reasons.)
func ReadAs[T any](u *ApiUrl) T {
	var v T
	u.ReadInto(&v)
	return v
}

// ------------------------------------------------------------
// Package-local

var repoUrl *ApiUrl

func doRequest(req *http.Request) *http.Response {
	if options.DryRun() {
		os.Exit(0)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	util.QuietlyHandle(err)
	return resp
}

func queryFromOptions() string {
	query := url.Values{}
	if stateVal := options.StateFlags().ToParam(); len(stateVal) > 0 {
		query.Set("state", stateVal)
	}

	// with an authenticated user, GitLab by default returns only
	// issues created by that user
	if accessToken := options.AccessToken(); len(accessToken) > 0 {
		query.Set("scope", "all")
	}

	if milestone := options.Milestone(); len(milestone) > 0 {
		var milestonesVal = url.QueryEscape(milestone)
		query.Set("milestone", milestonesVal)
	}

	// can't use query.Set() bc GitLab wants raw commas, not %2F
	queryString := query.Encode()
	if labels := options.Labels(); len(labels) > 0 {
		labelsVal := strings.Join(labels, ",")
		queryString = queryString + "&labels=" + labelsVal
	}

	return queryString
}
