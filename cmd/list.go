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
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(list())
	},
}

func list() error {
	// util.Log("Using token: ", config.Token)

	issuesUrl := urls.IssuesUrl()
	issuesUrl.RawQuery = params.ToRawQuery()

	util.Log(issuesUrl)

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
