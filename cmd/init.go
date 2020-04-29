package cmd

import "github.com/sirupsen/logrus"

func init() {
	if err := createPathIfNotExists(migrationsDest); err != nil {
		logrus.Fatalf("failed to create %s destination: %v", migrationsDest, err)
	}

	if err := createPathIfNotExists(cmdDest); err != nil {
		logrus.Fatalf("failed to create %s destination: %v", cmdDest, err)
	}

	Migration.AddCommand(Create)
	Migration.AddCommand(Template)
}
