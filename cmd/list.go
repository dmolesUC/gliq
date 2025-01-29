package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/apitypes"
	"github.com/dmolesUC/gliq/options"
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
	// TODO: pagination
	apiUrl := urls.IssuesUrl().WithParams()
	var issues = apitypes.ReadIssues(apiUrl)
	if options.Verbose() {
		util.Logf("Found %v issues\n", len(issues))
	}

	issues = filterByLinks(issues)
	if options.Verbose() {
		util.Logf("Filtered to %v issues\n", len(issues))
	}

	for _, issue := range issues {
		msg := issue.String()
		fmt.Println(msg)
	}
}

func filterByLinks(issues []apitypes.Issue) []apitypes.Issue {
	filtered := selectRelated(issues, options.RelatedIdsToInclude())
	return rejectRelated(filtered, options.RelatedIdsToExclude())
}

func selectRelated(issues []apitypes.Issue, relatedIds []int) []apitypes.Issue {
	if len(relatedIds) == 0 {
		return issues
	}
	//util.Logf("includeRelated: %v\n", relatedIds)
	return util.Select(issues, func(issue apitypes.Issue) bool {
		return issue.IsLinkedToAny(relatedIds)
	})
}

func rejectRelated(issues []apitypes.Issue, relatedIds []int) []apitypes.Issue {
	if len(relatedIds) == 0 {
		return issues
	}
	//util.Logf("excludeRelated: %v\n", relatedIds)
	return util.Reject(issues, func(issue apitypes.Issue) bool {
		return issue.IsLinkedToAny(relatedIds)
	})
}

func init() {
	rootCmd.AddCommand(listCmd)
}
