package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGitHubAccessToken() string {
	return githubAccessToken
}