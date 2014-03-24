package services

import (
	"github.com/mashup-cms/mashup-api/globals"
	"github.com/mashup-cms/mashup-api/model"
)

func FindGithubAdmins(login string, userId int) (*[]model.User, error) {
	users := []model.User{}
	adminQuery := "select users.id as id, users.name as name, users.email as email, users.login as login, " +
		"users.encrypted_password as encrypted_password, users.created_at as created_at, users.updated_at " +
		"as updated_at from users, memberships, github_accounts where github_accounts.login = :login and memberships.github_account_id = github_accounts.id " +
		"and users.id = memberships.user_id and memberships.is_admin"
	_, err := globals.PostgresConnection.Select(&users, adminQuery, map[string]interface{}{
		"login": login,
	})
	return &users, err
}
