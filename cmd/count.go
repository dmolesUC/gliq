package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

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

	resp, err := http.Get(apiUrl.String())
	cobra.CheckErr(err)

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
