package services

import (
	"github.com/mashup-cms/mashup-api/globals"
	"github.com/mashup-cms/mashup-api/model"

	"errors"
)

func FindGithubAccount(login string, userId int) (*model.GithubAccount, error) {
	account := &model.GithubAccount{}
	err := globals.PostgresConnection.SelectOne(account, "select * from github_accounts where login = :login", map[string]interface{}{
		"login": login,
	})
	if account.Id == 0 {
		err = errors.New("Account not found.")
	}
	return account, err
}
