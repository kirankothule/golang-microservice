package services

import (
	"net/http"
	"github.com/kirankothule/golang-microservice/mvc/domain"
	"github.com/kirankothule/golang-microservice/mvc/utils"
)

type itemService struct {}

var (
	ItemService itemService
)

func (i itemService) GetItem(item_id int) (*domain.Item, *utils.ApplicationErr) {
	return nil, &utils.ApplicationErr {
		Message: "implement me",
		StatusCode: http.StatusInternalServerError,

	}
}