package cmd

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	if err := createPathIfNotExists(destination); err != nil {
		logrus.Fatalf("failed to create %s destination: %v", destination, err)
	}

	createMigrationInterface()

	Migration.AddCommand(Create)
}

var destination = "migrations"

// Migration root command
var Migration = &cobra.Command{
	Use:   "",
	Short: "Migration tool",
	Long:  `Migration tool`,
}

func createPathIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func writeFileContent(name string, content []byte) {
	fName := destination + "/" + name

	if err := ioutil.WriteFile(fName, content, 0644); err != nil {
		logrus.Fatal("write content to file failed: ", err)
	}
}

func copy(from string, to string) error {
	b, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(to, b, 0644); err != nil {
		return err
	}

	return nil
}

func createMigrationInterface() {
	from := "cmd/migration.tpl"
	dest := destination + "/migration.go"

	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := copy(from, dest); err != nil {
			logrus.Fatal("failed to copy migration.tpl: ", err)
		}
	}
}
