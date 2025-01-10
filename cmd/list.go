package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/config"
	"github.com/dmolesUC/gliq/params"
	"github.com/dmolesUC/gliq/util"
)

const apiBaseUrlStr = "https://gitlab.com/api/v4"

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
	util.Log("Using token: ", config.Token)

	if config.Repo == "" {
		return fmt.Errorf("no repository specified")
	}
	repoUrlStr, err := url.JoinPath(apiBaseUrlStr, "projects", url.PathEscape(config.Repo))
	if err != nil {
		return err
	}

	issuesUrlStr, err := url.JoinPath(repoUrlStr, "issues")
	if err != nil {
		return err
	}

	issuesUrl, err := url.Parse(issuesUrlStr)
	if err != nil {
		return err
	}

	state, err := params.State()
	if err != nil {
		return err
	}

	q := issuesUrl.Query()
	q.Set("state", state)

	util.Log(issuesUrl)

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
