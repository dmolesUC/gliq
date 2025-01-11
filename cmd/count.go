package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/config"
	"github.com/dmolesUC/gliq/params"
	"github.com/dmolesUC/gliq/urls"
	"github.com/dmolesUC/gliq/util"
)

type Counts struct {
	All    int64
	Closed int64
	Opened int64
}

func (counts Counts) Included() int64 {
	var count int64
	switch params.StatesToInclude() {
	case params.Closed:
		count = counts.Closed
	case params.Opened:
		count = counts.Opened
	case params.All:
		count = counts.All
	}
	return count
}

type Statistics struct {
	Counts Counts
}

type StatsMessage struct {
	Statistics Statistics
}

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count issues",
	Long: util.Prettify(`
	Counts issues satisfying the specified criteria.
    `),
	Run: execCount,
}

func execCount(cmd *cobra.Command, args []string) {
	apiUrl := urls.IssueStatsUrl()
	apiUrl.RawQuery = params.ToRawQuery()

	util.Log(apiUrl)

	resp := Get(apiUrl)

	var stats = toStatistics(resp)

	counts := stats.Counts
	count := counts.Included()
	fmt.Println(count)
}

func toStatistics(resp *http.Response) Statistics {
	var msg StatsMessage

	defer util.QuietlyClose(resp.Body)
	err := json.NewDecoder(resp.Body).Decode(&msg)

	cobra.CheckErr(err)
	return msg.Statistics
}

func init() {
	rootCmd.AddCommand(countCmd)
}

// TODO: Move this to its own file
func Get(apiUrl *url.URL) *http.Response {
	req, err := http.NewRequest("get", apiUrl.String(), nil)
	cobra.CheckErr(err)

	if token := config.Token; len(token) > 0 {
		req.Header.Add("Authentication", "Authorization: Bearer "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	cobra.CheckErr(err)

	return resp
}
