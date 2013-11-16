package services

import (
  "github.com/mashup-cms/mashup-api/model"
  "github.com/mashup-cms/mashup-api/globals"
)

func FindUserById(userId int) (*model.User, error) {
  account, err := globals.PostgresConnection.Get(model.User{}, userId)
  if err != nil {
    return &model.User{}, err
  } else {
    return account.(*model.User), nil
  }
}