package model

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/mashup-cms/mashup-api/globals"
)

func FindByParams(modelArray interface{}, params map[string][]string, userId int) error {
	tableName := ""
	var obj interface{}
	whereClause := " where "
	andClause := " and "
	var args []interface{}
	var clauses []string
	currentVar := 1

	switch arrayType := modelArray.(type) {
	case *[]GithubAccount:
		tableName = "github_accounts"
		obj = GithubAccount{}
	case *[]Membership:
		tableName = "memberships"
		obj = Membership{}
	case *[]Repo:
		tableName = "repositories"
		obj = Repo{}
	case *[]User:
		tableName = "users"
		obj = User{}
	default:
		log.Printf("invalid %s", arrayType)
	}

	params = mapParams(obj, params)

	stmtBegin := fmt.Sprintf("select * from %s", tableName)
	buffer := bytes.NewBufferString(stmtBegin)

	for key, value := range params {

		var vals []string

		for i := currentVar; i < (currentVar + len(value)); i++ {
			vals = append(vals, fmt.Sprintf("$%v", i))
		}

		currentVar = currentVar + len(value)

		clause := fmt.Sprintf("%s in (%s)", key, strings.Join(vals, ", "))

		clauses = append(clauses, clause)

		for _, param := range value {
			args = append(args, param)
		}
	}

	if len(params) > 0 {
		buffer.WriteString(whereClause)
		buffer.WriteString(strings.Join(clauses, andClause))
	}

	_, err := globals.PostgresConnection.Select(modelArray, buffer.String(), args...)

	return err
}

func mapParams(obj interface{}, params map[string][]string) map[string][]string {
	structObj := reflect.TypeOf(obj)
	numFields := structObj.NumField()
	var fields map[string][]string
	fields = make(map[string][]string)
	for i := 0; i < numFields; i++ {
		field := structObj.Field(i)
		jsonName := strings.Split(field.Tag.Get("json"), ",")[0]
		if val, ok := params[jsonName]; ok {
			dbName := strings.Split(field.Tag.Get("db"), ",")[0]
			fields[dbName] = val
		}
	}
	return fields
}
