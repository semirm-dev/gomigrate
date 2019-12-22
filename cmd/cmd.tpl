package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Auto-generated file: https://github.com/semirm-dev/gomigrate
// Feel free to edit

func init() {
	Migration.AddCommand(Apply)
}

// Migration command
var Migration = &cobra.Command{
	Use:   "",
	Short: "Migrations tool",
	Long:  `Migrations tool`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Apply migrations
var Apply = &cobra.Command{
	Use:   "migrate",
	Short: "Apply migrations",
	Long:  `Apply migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Run()

		logrus.Info("migrations applied")
	},
}
