package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Template migration command
var Template = &cobra.Command{
	Use:   "tpl",
	Short: "Generate templates",
	Long:  `Generate templates`,
	Run: func(cmd *cobra.Command, args []string) {
		createPathIfNotExists(cmdDest)

		createMigrationInterface()

		createRegisterMigrationsCollection()

		createCmdApply()

		logrus.Info("templates generated")
	},
}
