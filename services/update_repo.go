package services

import (
  "github.com/mashup-cms/mashup-api/model"
  "github.com/mashup-cms/mashup-api/globals"
)

func UpdateRepo(repoId int, active bool, userId int) (*model.Repo, error) {
  obj, err := globals.PostgresConnection.Get(model.Repo{},repoId)
  if obj == nil {
    return &model.Repo{}, err
  }
  repo := obj.(*model.Repo)
  repo.Active = active
  _, err = globals.PostgresConnection.Update(repo)
  return repo, err
}