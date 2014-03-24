package services

import (
  "github.com/mashup-cms/mashup-api/model"
  "github.com/mashup-cms/mashup-api/globals"
  "github.com/mashup-cms/mashup-api/helpers"
  "github.com/mashup-cms/mashup-api/workers"
)

func SyncGithubAccount(taskId, login string, userId int, token string) (*model.Task, error) {
  task := &model.Task{taskId, "githubAccountSync"}
  account := &model.GithubAccount{}
  err := globals.PostgresConnection.SelectOne(account, "select * from github_accounts where login = :login", map[string]interface{}{
    "login": login,
  })
  if err == nil {
    //add queueing

    if (globals.ExternalQueue) {
      helpers.Enqueue("github_sync", "SynchronizeUser", []int{account.Id})
    } else {
      workers.NewTask(taskId, token, "SynchronizeUser", float64(account.Id))
  	}
    account.IsSyncing = true
    _, err := globals.PostgresConnection.Update(account)
    return task, err
  }
  return task, err
}