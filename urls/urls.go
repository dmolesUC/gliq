package urls

import (
	"net/url"

	"github.com/dmolesUC/gliq/config"
)

// ------------------------------------------------------------
// Exported functions

func RepoUrl() *ApiUrl {
	if repoUrl == nil {
		repo := url.PathEscape(config.Repo)
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

// ------------------------------------------------------------
// Private state

var repoUrl *ApiUrl
