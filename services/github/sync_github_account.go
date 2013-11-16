package github

import (
	"github.com/mashup-cms/mashup-api/globals"
	"github.com/mashup-cms/mashup-api/model"
	"code.google.com/p/goauth2/oauth"
	"database/sql"
	"github.com/google/go-github/github"
	"github.com/lib/pq"
	"log"
	"time"
)

func SyncGithubAccount(queue string, args ...interface{}) error {
	obj, err := globals.PostgresConnection.Get(model.GithubAccount{}, int(args[0].(float64)))
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	account := obj.(*model.GithubAccount)

	transport := &oauth.Transport{
		Token: &oauth.Token{AccessToken: account.GithubOauthToken},
	}

	client := github.NewClient(transport.Client())

	//Update Account Info
	err = updateAccount(client, account)
	if err != nil {
		return err
	}
	//Update Repositories
	err = updateRepositories(client, account)
	if err != nil {
		return err
	}
	//Update Syncing status
	account.SyncedAt = model.NullTime{pq.NullTime{time.Now().UTC(), true}}
	account.IsSyncing = false
	_, err = globals.PostgresConnection.Update(account)
	return err
}

func updateRepositories(client *github.Client, account *model.GithubAccount) error {
	repos, _, err := client.Repositories.List("", nil)
	if err != nil {
		return err
	}

	adminQuery := "select users.id as id, users.name as name, users.email as email, users.login as login, " +
		"users.encrypted_password as encrypted_password, users.created_at as created_at, users.updated_at " +
		"as updated_at from users, memberships where memberships.github_account_id = :id " +
		"and users.id = memberships.user_id and memberships.is_admin"
	admins := []model.User{}
	_, err = globals.PostgresConnection.Select(&admins, adminQuery, map[string]interface{}{
		"id": account.Id,
	})
	if err != nil {
		return err
	}

	for _, repo := range repos {
		obj, err := globals.PostgresConnection.Get(model.Repo{}, *repo.ID)
		if err != nil {
			return err
		}
		newRepo := &model.Repo{}
		if obj == nil {
			newRepo = &model.Repo{
				Id:            *repo.ID,
				Name:          *repo.Name,
				Owner:         *repo.Owner.Login,
				Active:        false,
				DefaultBranch: *repo.DefaultBranch,
			}

			if *repo.URL != "" {
				newRepo.Url = model.NullString{sql.NullString{*repo.URL, true}}
			}

			if *repo.Description != "" {
				newRepo.Description = model.NullString{sql.NullString{*repo.Description, true}}
			}

			err = globals.PostgresConnection.Insert(newRepo)
			if err != nil {
				return err
			}

		} else {
			newRepo = obj.(*model.Repo)
		}

		for _, admin := range admins {
			obj, err := globals.PostgresConnection.Get(model.Permission{}, newRepo.Id, admin.Id)
			if err != nil {
				return err
			}

			if obj != nil {
				perm := obj.(*model.Permission)
				perm.Admin = true
				_, err = globals.PostgresConnection.Update(perm)
				if err != nil {
					return err
				}

			} else {
				perm := &model.Permission{
					UserId: admin.Id,
					RepoId: newRepo.Id,
					Admin:  true,
				}
				err = globals.PostgresConnection.Insert(perm)
			}
		}
	}
	return err
}

func updateAccount(client *github.Client, account *model.GithubAccount) error {
	updated := false
	githubAccount, _, err := client.Users.Get("")
	if err != nil {
		return err
	}

	if *githubAccount.Login != account.Login {
		account.Login = *githubAccount.Login
		updated = true
	}
	if *githubAccount.GravatarID != account.GravatarId {
		account.GravatarId = *githubAccount.GravatarID
		updated = true
	}

	if updated {
		_, err = globals.PostgresConnection.Update(account)
		if err != nil {
			return err
		}
	}
	return nil
}
