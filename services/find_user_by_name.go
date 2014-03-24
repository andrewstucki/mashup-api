package services

import (
	"github.com/mashup-cms/mashup-api/globals"
	"github.com/mashup-cms/mashup-api/model"
	
	"log"
)

func FindUserByName(login string, userId int) (*model.User, error) {
	account := &model.User{}
	log.Printf("login %s",login)
	err := globals.PostgresConnection.SelectOne(account, "select * from users where login = :login", map[string]interface{}{
		"login": login,
	})
	return account, err
}
