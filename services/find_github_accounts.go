package services

import (
	"github.com/mashup-cms/mashup-api/model"
)

func FindGithubAccounts(params map[string][]string, userId int) (*model.GithubAccounts, error) {
	accounts := model.GithubAccounts{Accounts: []model.GithubAccount{}}
	err := model.FindByParams(&accounts.Accounts, params, userId)
	return &accounts, err
}
