package model

import (
	"time"
)

type GithubAccounts struct {
	Accounts []GithubAccount `json:"githubAccounts"`
}

type GithubAccountService struct {
	Account *GithubAccount `json:"githubAccount"`
}

type GithubAccount struct {
	Id               int       `db:"id" json:"id"`
	Login            string    `db:"login" json:"login"`
	IsSyncing        bool      `db:"is_syncing" json:"isSyncing"`
	SyncedAt         NullTime  `db:"synced_at" json:"syncedAt"`
	GithubOauthToken string    `db:"github_oauth_token" json:"-"`
	GravatarId       string    `db:"gravatar_id" json:"gravatarId"`
	CreatedAt        time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt        time.Time `db:"updated_at" json:"updatedAt"`
}

func (account *GithubAccount) PreInsert(interface{}) error {
	account.CreatedAt = time.Now().UTC()
	account.UpdatedAt = account.CreatedAt
	return nil
}

func (account *GithubAccount) PreUpdate(interface{}) error {
	account.UpdatedAt = time.Now().UTC()
	return nil
}

func (account *GithubAccount) GetService() interface{} {
	return GithubAccountService{Account: account}
}
