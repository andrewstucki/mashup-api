package services

import (
  "github.com/mashup-cms/mashup-api/model"
  "github.com/mashup-cms/mashup-api/globals"
)

func AddGithubAdmins(login string, users *model.Users, userId int) (error) {
  account, err := FindGithubAccount(login, userId)
  if err == nil {
    for _, user := range users.Users {
      membership := &model.Membership {
        GithubAccountId: account.Id,
        UserId: user.Id,
        IsAdmin: true,
      }
      err = globals.PostgresConnection.Insert(membership)      
    }
  }
  return err
}