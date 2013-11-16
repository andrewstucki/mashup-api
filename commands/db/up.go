package db

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"log"
	
	"github.com/mashup-cms/mashup-api/commands"
)

var UpCmd = &commands.Command{
	Name:    "db:up",
	Usage:   "",
	Summary: "Migrate the DB to the most recent version available",
	Help:    `up extended help here...`,
	Run:     upRun,
}

func upRun(cmd *commands.Command, args ...string) {

	conf, err := dbConfFromFlags()
	if err != nil {
		log.Fatal(err)
	}

	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := goose.RunMigrations(conf, conf.MigrationsDir, target); err != nil {
		log.Fatal(err)
	}
}
