package main

import (
	"github.com/semirm-dev/gomigrate/cmd"
)

func main() {
	cmd.Create.Flags().StringP("migration", "m", "", "migration name")
	cmd.Create.MarkFlagRequired("migration")

	cmd.Migration.Execute()
}
