package services

import (
	"github.com/mashup-cms/mashup-api/globals"
	"github.com/mashup-cms/mashup-api/model"

	"strconv"
)

func FindUser(id string, userId int) (*model.User, error) {
	idNum, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return &model.User{}, err
	}
	idN := int(idNum)
	obj, err := globals.PostgresConnection.Get(model.User{}, idN)
	var account *model.User
	if err == nil {
		account = obj.(*model.User)
	}
	return account, err
}
