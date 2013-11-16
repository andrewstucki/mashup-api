package model

type Membership struct {
	GithubAccountId int  `db:"github_account_id"`
	UserId          int  `db:"user_id"`
	IsAdmin         bool `db:"is_admin"`
}
