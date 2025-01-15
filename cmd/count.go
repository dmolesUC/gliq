package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/messages"
	"github.com/dmolesUC/gliq/params"
	"github.com/dmolesUC/gliq/urls"
	"github.com/dmolesUC/gliq/util"
)

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

	statsMessage := urls.ReadFromUrl[statsMessage](apiUrl)
	stats := statsMessage.Statistics

	counts := stats.Counts
	count := counts.Included()
	fmt.Println(count)
}

func init() {
	rootCmd.AddCommand(countCmd)
}

type statsMessage struct {
	Statistics messages.Statistics
}
