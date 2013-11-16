package model

import (
	"time"

	"github.com/mashup-cms/mashup-api/globals"
)

type Repos struct {
	Repos []Repo `json:"repos"`
}

type RepoService struct {
	Repo *Repo `json:"repo"`
}

type Repo struct {
	Id            int        `db:"id" json:"id"`
	Name          string     `db:"name" json:"name"`
	Url           NullString `db:"url" json:"url,omitempty"`
	Owner         string     `db:"owner_name" json:"owner"`
	Active        bool       `db:"active" json:"active"`
	Description   NullString `db:"description" json:"description"`
	DefaultBranch string     `db:"default_branch" json:"defaultBranch"`
	VimeoKey      NullString `db:"vimeo_key" json:"-"`
	FlickrKey     NullString `db:"flickr_key" json:"-"`
	CreatedAt     time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updatedAt"`
}

func (repo *Repo) Activate() error {
	repo.Active = !repo.Active
	_, err := globals.PostgresConnection.Update(repo)
	return err
}

func (repo *Repo) PreInsert(interface{}) error {
	repo.CreatedAt = time.Now().UTC()
	repo.UpdatedAt = repo.CreatedAt
	return nil
}

func (repo *Repo) PreUpdate(interface{}) error {
	repo.UpdatedAt = time.Now().UTC()
	return nil
}

func (repo *Repo) GetService() interface{} {
	return RepoService{Repo: repo}
}
