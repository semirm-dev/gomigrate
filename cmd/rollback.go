package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Rollback migrations command
var Rollback = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback migrations",
	Long:  `Rollback migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("migrations rolled back")
	},
}
