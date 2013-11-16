package services

import (
  "log"

  "github.com/google/go-github/github"
  "code.google.com/p/goauth2/oauth"
  
  "github.com/mashup-cms/mashup-api/model"
  "github.com/mashup-cms/mashup-api/globals"
)

func AddGithubToken(token string, userId int) (*model.GithubAccount, error) {
  //error checking isn't real right now
  var account *model.GithubAccount
  transport := &oauth.Transport{
    Token: &oauth.Token{AccessToken: token},
  }
  
  client := github.NewClient(transport.Client())
  githubAccount, data, err := client.Users.Get("")
  if err != nil {
    log.Printf("%s, %s, %s",err.Error(), data, transport.Token)
  }
  if err == nil {
    updated := false
    obj, _ := globals.PostgresConnection.Get(model.GithubAccount{}, *githubAccount.ID)
    if obj == nil {
      //create
      account = &model.GithubAccount{
        Id: *githubAccount.ID,
        Login: *githubAccount.Login,
        IsSyncing: false,
        GithubOauthToken: token,
        GravatarId: *githubAccount.GravatarID,
      }
      err := globals.PostgresConnection.Insert(account)
      if err == nil {
        updated = true
      }
    } else {
      //bind and update names
      account = obj.(*model.GithubAccount)
      if account.Login != *githubAccount.Login {
        query := "update repositories set owner_name=? where owner_name=?"
        trans, err := globals.PostgresConnection.Begin()
        if err == nil {
          trans.Exec(query, *githubAccount.Login, account.Login)
          account.Login = *githubAccount.Login
          trans.Update(&account)
          err := trans.Commit()
          if err == nil {
            updated = true
          }
        }
      } else {
        updated = true
      }
    }
    if updated {
      obj, _ := globals.PostgresConnection.Get(model.Membership{}, account.Id, userId)
      if obj == nil {
        membership := &model.Membership{
          GithubAccountId: account.Id,
          UserId: userId,
          IsAdmin: true,
        }
        err := globals.PostgresConnection.Insert(membership)
        if err == nil {
          return account, nil
        }
      } else {
        return account, nil
      }
    }
    
    //error on insert/update
    log.Printf("error")
    return account, nil
  } else {
    log.Printf("no user found")
    return account, nil
  }  
}