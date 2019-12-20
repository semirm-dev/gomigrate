package cmd

import (
	"fmt"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type migration struct {
	Name      string
	Timestamp int64
	Applied   bool
}

// Apply migrations command
var Apply = &cobra.Command{
	Use:   "apply",
	Short: "Apply migrations",
	Long:  `Apply migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		// get migrations from database
		migrations := []*migration{
			&migration{
				Name:      "migration3",
				Timestamp: 1576855900,
				Applied:   false,
			},
			&migration{
				Name:      "migration1",
				Timestamp: 1576855882,
				Applied:   false,
			},
			&migration{
				Name:      "migration2",
				Timestamp: 1576855897,
				Applied:   true,
			},
			&migration{
				Name:      "migration4",
				Timestamp: 1576855934,
				Applied:   false,
			},
		}

		sort.Slice(migrations, func(i, j int) bool {
			return migrations[i].Timestamp < migrations[j].Timestamp
		})

		for _, m := range migrations {
			if !m.Applied {
				fName := "Apply" + fmt.Sprint(m.Timestamp)

				// TODO
				logrus.Info(fName)
			}
		}

		logrus.Info("migrations applied")
	},
}
