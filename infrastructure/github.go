package infrastructure

import (
	"net/http"

	"github.com/google/go-github/v24/github"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// GitHubClientFactory returns a function to create GitHub v3/v4 clients.
func GitHubClientFactory(token string) func() (*github.Client, *githubv4.Client) {
	return func() (*github.Client, *githubv4.Client) {
		hc := &http.Client{
			Transport: &oauth2.Transport{
				Base:   http.DefaultTransport,
				Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
			},
		}
		return github.NewClient(hc), githubv4.NewClient(hc)
	}
}