package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Create migration command
var Create = &cobra.Command{
	Use:   "create",
	Short: "Create migration",
	Long:  `Create migration`,
	Run: func(cmd *cobra.Command, args []string) {
		createMigrationInterface()

		var name = cmd.Flag("migration").Value.String()

		if strings.TrimSpace(name) == "" {
			logrus.Fatal("invalid migration name")
		}

		createMigration(name)

		logrus.Info("migration created")
	},
}
