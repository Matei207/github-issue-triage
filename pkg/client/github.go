package client

import (
	"encoding/json"
	"github.com/matei207/github-issue-triage/pkg/api"
	"github.com/matei207/github-issue-triage/pkg/helper"
	"github.com/matei207/github-issue-triage/pkg/options"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

var logger *zap.SugaredLogger

const (
	GITHUB_URL = "https://api.github.com/repos/linkerd/linkerd2/issues"
)

func init() {
	logger = helper.NewZapSugaredLogger()
}

type GitHubClient struct {
	c     *http.Client
	token string
}

func NewGitHubClient(token string) *GitHubClient {
	return &GitHubClient{
		c:     &http.Client{},
		token: token,
	}
}

func (gh *GitHubClient) Do(r *http.Request) (*http.Response, error) {
	return gh.c.Do(r)
}

func Run(opt *options.Options) {
	logger.Infof("Sending request to GitHub:%s", GITHUB_URL)
	clientToken := "token " + opt.ApiToken
	gh := NewGitHubClient(clientToken)

	req, err := http.NewRequest("GET", GITHUB_URL, nil)
	if err != nil {
		logger.Errorw("error creating req", "url", GITHUB_URL, "err", err)
		return
	}

	req.Header.Add("Authorization", clientToken)
	resp, err := gh.Do(req)
	if err != nil {
		logger.Errorw("error sending request", "err", err)
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var issues []*api.GitHubIssue

	if err := json.Unmarshal(data, &issues); err != nil {
		logger.Errorw("error unmarshalling json data", "payload", data, "err", err)
	}

	for i, issue := range issues {
		logger.Infof("Issue #%d: <Name:%s -- Status:%s -- Created:%v>", i, issue.Title, issue.State, issue.TsCreated)
	}
}
