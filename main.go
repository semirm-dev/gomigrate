package main

import (
	"github.com/semirm-dev/gomigrate/cmd"
)

var migrationName string

func main() {
	cmd.Create.Flags().StringVar(&migrationName, "migration", "", "migration name")
	cmd.Create.MarkFlagRequired("migration")

	cmd.Migration.Execute()
}
