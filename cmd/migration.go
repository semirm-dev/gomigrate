package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/dave/jennifer/jen"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	if err := createPathIfNotExists(destination); err != nil {
		logrus.Fatal("failed to create migration.go interface: ", err)
	}

	Migration.AddCommand(Create)
	Migration.AddCommand(Verify)
	Migration.AddCommand(Apply)
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

func createMigration(name string) {
	m := jen.NewFile("migrations")

	m.Comment(name + " migration")
	m.Type().Id(name).Struct()

	m.Comment("Apply migration")
	m.Func().Params(jen.Id("mig").Id("*" + name)).Id("Apply").Params().Block(
		jen.Qual("github.com/sirupsen/logrus", "Info").Call(jen.Lit("Applying migration")),
	)

	ts := time.Now().Unix()

	m.Comment("Timestamp when migration was created")
	m.Func().Params(jen.Id("mig").Id("*" + name)).Id("Timestamp").Params().Int64().Block(
		jen.Return(jen.Lit(ts)),
	)

	n := strings.ToLower(fmt.Sprint(ts) + "_" + name + ".go")
	c := []byte(fmt.Sprintf("%#v", m))

	writeFileContent(n, c)
}

func createMigrationInterface() {
	if _, err := os.Stat(destination + "/migration.go"); os.IsNotExist(err) {
		m := jen.NewFile("migrations")

		m.Comment("Collection with all migrations")
		m.Var().Id("Collection").Op("=").Make(jen.Index().Id("Migration").Op(",").Lit(0))

		m.Comment("Migration service")
		m.Type().Id("Migration").Interface(jen.Id("Apply").Params(), jen.Id("TimeStamp").Params().Int64())

		c := []byte(fmt.Sprintf("%#v", m))

		writeFileContent("migration.go", c)
	}
}
