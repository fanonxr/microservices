package services

import (
	"microservices/mvc/domain"
	"microservices/mvc/utils"
	"net/http"
)

// setting this up allows no public functions to be added to same packaged
// but instead tied to the interface

type itemService struct {
	
}

var (
	ItemsService itemService
)

func (*itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError){
	return nil, &utils.ApplicationError{
		Message:    "",
		StatusCode: http.StatusInternalServerError,
		Code:       "",
	}
}
