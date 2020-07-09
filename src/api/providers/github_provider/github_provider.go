package github_provider
// "fbfe1fff05cf4568d230303927449d48e304a61"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"microservices/src/api/domain/github"
	"microservices/src/api/restclient"
	"net/http"
)

// providers is where we are actually making the request

const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization,getAuthorizationHeader(accessToken))
	
	res, err := restclient.Post(urlCreateRepo, request, headers)

	if err != nil {
		log.Print(fmt.Sprintf("Error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          err.Error(),
			DocumentationURL: "",
			Errors:           nil,
		}
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message:"invalid response body"}
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		var errRes github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errRes); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message:"invalid json response body"}
		}
		errRes.StatusCode = res.StatusCode
		return nil, &errRes
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo succesful response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode:http.StatusInternalServerError, Message:"error when trying to unmarshal"}
	}
	return &result, nil
}