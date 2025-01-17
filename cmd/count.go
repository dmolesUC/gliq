package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/apitypes"
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

func execCount(*cobra.Command, []string) {
	apiUrl := urls.IssueStatsUrl().WithParams()
	var statsResp = urls.ReadAs[apitypes.StatisticsResponse](apiUrl)

	counts := statsResp.Statistics.Counts
	fmt.Println(counts.Included())
}

func init() {
	rootCmd.AddCommand(countCmd)
}
