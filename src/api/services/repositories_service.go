package services

import (
	"microservices/src/api/config"
	"microservices/src/api/domain/github"
	"microservices/src/api/domain/repositories"
	"microservices/src/api/providers/github_provider"
	"microservices/src/api/utils/errors"
	"strings"
)

type repoService struct {

}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repository Name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: "",
		Homepage:    "",
		Private:     false,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGitHubAccessToken(), request)

	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	
	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}

	return &result, nil

}
