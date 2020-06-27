package services

import (
	"microservices/mvc/domain"
	"microservices/mvc/utils"
)
// defining struct for userService
type userService struct {

}

var (
	UserService userService
)
// having this implement the userSerivce struct
func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}