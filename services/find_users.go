package services

import (
	"github.com/mashup-cms/mashup-api/model"
)

func FindUsers(params map[string][]string, userId int) (*[]model.User, error) {
	users := []model.User{}
	err := model.FindByParams(&users, params, userId)
	return &users, err
}
