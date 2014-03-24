package services

import (
	"github.com/mashup-cms/mashup-api/model"
	"errors"
)

func FindRepos(params map[string][]string, userId int) (*[]model.Repo, error) {
	repos := []model.Repo{}
	err := model.FindByParams(&repos, params, userId)
	if len(repos) == 0 && err == nil {
		err = errors.New("No repository found.")
	}
	return &repos, nil
}
