package model

type Permission struct {
	UserId int  `db:"user_id"`
	RepoId int  `db:"repository_id"`
	Admin  bool `db:"admin"`
}
