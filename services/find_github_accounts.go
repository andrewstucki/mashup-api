package services

import (
	"github.com/mashup-cms/mashup-api/model"
)

func FindGithubAccounts(params map[string][]string, userId int) (*[]model.GithubAccount, error) {
	accounts := []model.GithubAccount{}
	err := model.FindByParams(&accounts, params, userId)
	return &accounts, err
}
