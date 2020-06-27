package domain

import (
	"fmt"
	"log"
	"microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User {
		123:{Id: 123, FirstName:"Frank", LastName:"Rogers", Email:"frank@gmail.com"},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

// creating an interface
type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {

}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("We're accessing the database")
	if user := users[userId]; user != nil{
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("use %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:       "Not Found",
	}
}