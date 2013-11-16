package connections

import (
	"github.com/mashup-cms/mashup-api/globals"
	"github.com/mashup-cms/mashup-api/model"
	
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"log"

	//"os"
)

func SetupPostgres() {
	db, err := sql.Open("postgres", "dbname=mashup-dev sslmode=disable")
	if err != nil {
		log.Fatal("Unable to connect to postgres: %v", err)
	} else {
		globals.PostgresConnection = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
		globals.PostgresConnection.AddTableWithName(model.User{}, "users").SetKeys(true, "Id")
		globals.PostgresConnection.AddTableWithName(model.GithubAccount{}, "github_accounts").SetKeys(false, "Id")
		globals.PostgresConnection.AddTableWithName(model.Membership{}, "memberships").SetKeys(false, "GithubAccountId", "UserId")
		globals.PostgresConnection.AddTableWithName(model.Permission{}, "permissions").SetKeys(false, "RepoId", "UserId")
		globals.PostgresConnection.AddTableWithName(model.Repo{}, "repositories").SetKeys(false, "Id")

		//globals.PostgresConnection.TraceOn("[gorp]", log.New(os.Stdout, "mashup:", log.Lmicroseconds))
	}
	return
}
