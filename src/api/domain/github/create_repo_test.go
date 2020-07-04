package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "goalng lesson rest repo",
		Description: "a golang lesson repo for testing purposes",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	if request.Private {

	}

	// Marshall takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(request)

	// testing to
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))
	assert.EqualValues(t, `{"name":"goalng lesson rest repo","description":"a golang lesson repo for testing purposes","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`, string(bytes))
}