package services

import (
	"golang/microservices/domain"
	"golang/microservices/utils"
	"net/http"
)

type itemService struct{}

var ItemService itemService

func (*itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Needs implementing",
		StatusCode: http.StatusInternalServerError,
	}
}
