package urls

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/config"
	"github.com/dmolesUC/gliq/util"
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

func Get(apiUrl *url.URL) *http.Response {
	urlStr := apiUrl.String()
	log.Printf("GET %#v", urlStr)

	req, err := http.NewRequest("GET", apiUrl.String(), nil)
	cobra.CheckErr(err)

	if token := config.Token; len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	cobra.CheckErr(err)

	return resp
}

func ReadFromUrl[T any](apiUrl *url.URL) T {
	var msg T

	resp := Get(apiUrl)

	defer util.QuietlyClose(resp.Body)

	body, err := io.ReadAll(resp.Body)
	cobra.CheckErr(err)

	bodyReader := bytes.NewReader(body)

	err = json.NewDecoder(bodyReader).Decode(&msg)
	if err != nil {
		log.Printf("verbose error info: %#v", err)
		log.Printf("body was: %#v\n", string(body[:]))
		cobra.CheckErr(err) // TODO: something smarter
	}
	return msg
}

func init() {
	apiBaseUrl, _ = url.Parse(apiBaseUrlStr)
}
