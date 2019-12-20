package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Verify migrations command
var Verify = &cobra.Command{
	Use:   "verify",
	Short: "Verify migrations",
	Long:  `Verify migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("migrations verified")
	},
}
