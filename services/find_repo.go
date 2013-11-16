package services

import (
  "github.com/mashup-cms/mashup-api/model"
  "github.com/mashup-cms/mashup-api/globals"
)

func FindRepo(account string, name string, userId int) (*model.Repo, error) {
  repo := &model.Repo{}
  err := globals.PostgresConnection.SelectOne(repo, "select * from repositories where owner_name = :account and name = :name", map[string]interface{}{
    "account": account,
    "name": name,
  })
  return repo, err
}