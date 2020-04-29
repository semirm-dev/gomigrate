package cmd

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func init() {
	createRequiredDirectories()

	parseConfigFile()

	Migration.AddCommand(Create)
	Migration.AddCommand(Template)
}

func createRequiredDirectories() {
	if err := createPathIfNotExists(migrationsDest); err != nil {
		logrus.Fatalf("failed to create %s destination: %v", migrationsDest, err)
	}

	if err := createPathIfNotExists(cmdDest); err != nil {
		logrus.Fatalf("failed to create %s destination: %v", cmdDest, err)
	}
}

func parseConfigFile() {
	configYml, err := ioutil.ReadFile(cmdDest + "/config.yml")
	if err != nil {
		logrus.Fatalf("failed to read config.yml: %v", err)
	}

	err = yaml.Unmarshal(configYml, &Conf)
	if err != nil {
		logrus.Fatalf("failed to unmarshal config.yml: %v", err)
	}
}
