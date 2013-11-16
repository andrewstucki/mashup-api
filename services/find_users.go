package services

import (
	"github.com/mashup-cms/mashup-api/model"
)

func FindUsers(params map[string][]string, userId int) (*model.Users, error) {
	users := model.Users{Users: []model.User{}}
	err := model.FindByParams(&users.Users, params, userId)
	return &users, err
}
