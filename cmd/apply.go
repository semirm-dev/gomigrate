package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Apply migrations command
var Apply = &cobra.Command{
	Use:   "apply",
	Short: "Apply migrations",
	Long:  `Apply migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("migrations applied")
	},
}
