package services

import (
  "github.com/mashup-cms/mashup-api/model"
  "github.com/mashup-cms/mashup-api/globals"
  "github.com/mashup-cms/mashup-api/helpers"
)

func SyncGithubAccount(login string, userId int) (*model.GithubAccount, error) {
  account := &model.GithubAccount{}
  err := globals.PostgresConnection.SelectOne(account, "select * from github_accounts where login = :login", map[string]interface{}{
    "login": login,
  })
  if err == nil {
    //add queueing
    helpers.Enqueue("github_sync", "SynchronizeUser", []int{account.Id})
    account.IsSyncing = true
    _, err := globals.PostgresConnection.Update(account)
    return account, err
  }
  return account, err
}