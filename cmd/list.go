package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"

	"github.com/dmolesUC/gliq/messages"
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
	apiUrl := urls.IssuesUrl()
	apiUrl.RawQuery = params.ToRawQuery()

	resp := urls.Get(apiUrl)
	issues := toIssues(resp) // TODO: pagination
	for _, issue := range issues {
		msg := fmt.Sprintf("%v: %v", issue.Iid, issue.Title)
		if len(issue.Assignees) > 0 {
			assigneeUsernames := strings.Join(util.Map(issue.Assignees, (*messages.User).String), ", ")
			msg += fmt.Sprintf("(%v)", assigneeUsernames)
		}
		fmt.Println(msg)
	}
}

func toIssues(resp *http.Response) []messages.Issue {
	var msg []messages.Issue

	defer util.QuietlyClose(resp.Body)

	// TODO: share this code
	body, err := io.ReadAll(resp.Body)
	cobra.CheckErr(err)

	bodyReader := bytes.NewReader(body)

	err = json.NewDecoder(bodyReader).Decode(&msg)
	if err != nil {
		log.Printf("verbose error info: %#v\n", err)
		log.Printf("body was: %#v\n", string(body[:]))
	}

	cobra.CheckErr(err)

	return msg
}

func init() {
	rootCmd.AddCommand(listCmd)
}
