package main

import (
	"github.com/smira/goose/lib/goose"
	"log"
)

var upAllCmd = &Command{
	Name:    "up_all",
	Usage:   "",
	Summary: "Migrate the DB to the most recent version available, including ALL pending migrations as well",
	Help:    `up_all extended help here...`,
	Run:     upAllRun,
}

func upAllRun(cmd *Command, args ...string) {

	conf, err := dbConfFromFlags()
	if err != nil {
		log.Fatal(err)
	}

	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := goose.RunPendingMigrations(conf, conf.MigrationsDir, target); err != nil {
		log.Fatal(err)
	}
}
