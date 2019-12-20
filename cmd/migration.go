package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	Migration.AddCommand(Create)
	Migration.AddCommand(Verify)
	Migration.AddCommand(Apply)
	Migration.AddCommand(Rollback)
}

// Migration root command
var Migration = &cobra.Command{
	Use:   "",
	Short: "Migration tool",
	Long:  `Migration tool`,
}
