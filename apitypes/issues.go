package apitypes

import (
	"fmt"
	"net/url"
	"slices"
	"strings"

	"github.com/dmolesUC/gliq/urls"
	"github.com/dmolesUC/gliq/util"
)

// ------------------------------------------------------------
// Exported

type Issue struct {
	// ----------------------------------------
	// Exported, unmarshaled from JSON

	Iid         int64
	Title       string
	Description string
	State       string
	Labels      []string
	Milestone   *Issue
	Author      *User
	Assignees   []*User
	WebUrl      string `json:"web_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`

	// ----------------------------------------
	// Package-local, lazily initialized

	issueUrl       *urls.ApiUrl
	linksUrl       *urls.ApiUrl
	linkedIssues   []Issue
	linkedIssueIds []int // TODO: standardize int vs int64
}

func ReadIssues(apiUrl *urls.ApiUrl) []Issue {
	return urls.ReadAs[[]Issue](apiUrl)
}

func (i Issue) String() string {
	str := fmt.Sprintf("%v: %v", i.Iid, i.Title)
	if len(i.Assignees) > 0 {
		assigneeNames := util.Map(i.Assignees, (*User).String)
		str += fmt.Sprintf("(%v)", strings.Join(assigneeNames, ", "))
	}
	if linkedIssues := i.LinkedIssues(); len(linkedIssues) > 0 {
		linkedIssueIids := util.Map(linkedIssues, func(i2 Issue) string {
			return fmt.Sprintf("%v", i2.Iid)
		})
		str += fmt.Sprintf(" (related: %v)", strings.Join(linkedIssueIids, ", "))
	}
	return str
}

func (i Issue) Url() *urls.ApiUrl {
	if i.issueUrl == nil {
		u, _ := url.Parse(i.WebUrl)
		i.issueUrl = &urls.ApiUrl{URL: *u}
	}
	return i.issueUrl
}

func (i Issue) LinksUrl() *urls.ApiUrl {
	if i.linksUrl == nil {
		i.linksUrl = i.Url().JoinPath("links")
	}
	return i.linksUrl
}

func (i Issue) LinkedIssues() []Issue {
	if i.linkedIssues == nil {
		i.linkedIssues = urls.ReadAs[[]Issue](i.LinksUrl())
	}
	return i.linkedIssues
}

func (i Issue) IsLinkedToAny(relatedIds []int) bool {
	return slices.ContainsFunc(relatedIds, i.isLinkedTo)
}

// ------------------------------------------------------------
// Package-local

func (i Issue) iidAsInt() int {
	return int(i.Iid)
}

func (i Issue) isLinkedTo(relatedId int) bool {
	var linkedIssueIds = i.linkedIssueIds
	if linkedIssueIds == nil {
		linkedIssueIds = util.Map(i.LinkedIssues(), Issue.iidAsInt)
		slices.Sort(linkedIssueIds)
		i.linkedIssueIds = linkedIssueIds
	}
	_, found := slices.BinarySearch(linkedIssueIds, relatedId)
	return found
}
