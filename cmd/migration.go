package cmd

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	if err := createPathIfNotExists(migrationsDest); err != nil {
		logrus.Fatalf("failed to create %s destination: %v", migrationsDest, err)
	}

	Migration.AddCommand(Create)
	Migration.AddCommand(Template)
}

var (
	migrationsDest = "migrations"
	cmdDest        = "cmd"
)

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

func writeFileContent(path string, content []byte) {
	if err := ioutil.WriteFile(path, content, 0644); err != nil {
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

func downloadTpl(path string, dest string) error {
	resp, err := http.Get(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
