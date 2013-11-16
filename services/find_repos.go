package services

import (
	"github.com/mashup-cms/mashup-api/model"
	"errors"
)

func FindRepos(params map[string][]string, userId int) (*model.Repos, error) {
	repos := model.Repos{Repos: []model.Repo{}}
	err := model.FindByParams(&repos.Repos, params, userId)
	if len(repos.Repos) == 0 && err == nil {
		err = errors.New("No repository found.")
	}
	return &repos, err
}
