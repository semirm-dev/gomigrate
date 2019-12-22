package main

import (
	"github.com/semirm-dev/gomigrate/cmd"
)

func main() {
	cmd.Create.Flags().StringP("migration", "m", "", "migration name")
	cmd.Create.MarkFlagRequired("migration")

	cmd.Template.Flags().StringP("pkg", "p", "", "github project url [my-github-username/my-project]")
	cmd.Template.MarkFlagRequired("pkg")

	cmd.Migration.Execute()
}
