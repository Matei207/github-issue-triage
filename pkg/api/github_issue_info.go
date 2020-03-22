package api

import "time"

type GitHubID int
type IssueState string

const (
	IssueOpen   IssueState = "open"
	IssueClosed IssueState = "closed"
)

type GitHubIssue struct {
	Id        GitHubID         `json:"id"`
	Assignee  GitHubAssignee   `json:"assignee,omitempty"`
	Assignees []GitHubAssignee `json:"assignees,omitempty"`
	Body      string           `json:"body"`
	TsCreated time.Time        `json:"created_at"`
	TsClosed  time.Time        `json:"closed_at,omitempty"`
	Locked    bool             `json:"locked"`
	State     IssueState       `json:"state"`
	Title     string           `json:"title"`
}

type GitHubAssignee struct {
	Id   GitHubID `json:"id"`
	Name string   `json:"login,omitempty"`
}
