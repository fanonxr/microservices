package services

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"microservices/src/api/domain/repositories"
	"microservices/src/api/restclient"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	restclient.StartMockUps()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	res, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "Invalid repository Name", err.Message())

}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/"}`)),
		},
		Err: nil,
	})

	request := repositories.CreateRepoRequest{Name: "testing"}
	res, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
		Err: nil,
	})

	request := repositories.CreateRepoRequest{Name: "testing"}
	res, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.EqualValues(t, 123, res.Id)
}
