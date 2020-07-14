package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS TOKEN", githubAccessToken)
}

func TestGetGitHubAccessToken(t *testing.T) {
	assert.EqualValues(t, "", GetGitHubAccessToken)
}
