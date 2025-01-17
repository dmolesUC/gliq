package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/apitypes"
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

func execList(*cobra.Command, []string) {
	apiUrl := urls.IssuesUrl().WithParams()
	var issues = urls.ReadAs[[]apitypes.Issue](apiUrl)
	log.Printf("Found %v issues\n", len(issues))

	for _, issue := range issues {
		msg := fmt.Sprintf("%v: %v", issue.Iid, issue.Title)
		if len(issue.Assignees) > 0 {
			assigneeNames := util.Map(issue.Assignees, (*apitypes.User).String)
			msg += fmt.Sprintf("(%v)", strings.Join(assigneeNames, ", "))
		}
		fmt.Println(msg)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
