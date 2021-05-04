package domain

import (
	"fmt"
	"golang/microservices/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Abdillah", LastName: "Abu Musa", Email: "email@email.com"},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("Accessing the DB")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "Not Found",
	}

}
