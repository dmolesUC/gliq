package urls

import (
	"net/url"

	"github.com/dmolesUC/gliq/config"
)

const apiBaseUrlStr = "https://gitlab.com/api/v4"

var apiBaseUrl *url.URL // TODO: move to... something global?
var repoUrl *url.URL    // TODO: move to Config?

func RepoUrl() *url.URL {
	if repoUrl == nil {
		repoUrl = apiBaseUrl.JoinPath("projects", url.PathEscape(config.Repo))
	}
	return repoUrl
}

func IssuesUrl() *url.URL {
	return RepoUrl().JoinPath("issues")
}

func IssueStatsUrl() *url.URL {
	return RepoUrl().JoinPath("issues_statistics")
}

func init() {
	apiBaseUrl, _ = url.Parse(apiBaseUrlStr)
}
