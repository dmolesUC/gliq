package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/options"
	"github.com/dmolesUC/gliq/util"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gliq",
	Short: "A GitLab issue query tool",
	Long: util.Prettify(`
	gliq is a tool for querying GitLab issues.
	`),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false

	options.InitOptions(rootCmd)
}
