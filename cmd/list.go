package cmd

import (
	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/params"
	"github.com/dmolesUC/gliq/urls"
	"github.com/dmolesUC/gliq/util"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List issues",
	Long: util.Prettify(`
	Lists issues satisfying the specified criteria.
	`),
	Run: execList,
}

func execList(cmd *cobra.Command, args []string) {
	// util.Log("Using token: ", config.Token)

	apiUrl := urls.IssuesUrl()
	apiUrl.RawQuery = params.ToRawQuery()

	// TODO: make the query & parse the response

	util.Log(apiUrl)
}

func init() {
	rootCmd.AddCommand(listCmd)
}
