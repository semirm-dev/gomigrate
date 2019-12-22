package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/dave/jennifer/jen"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Create migration command
var Create = &cobra.Command{
	Use:   "create",
	Short: "Create migration",
	Long:  `Create migration`,
	Run: func(cmd *cobra.Command, args []string) {
		var name = cmd.Flag("migration").Value.String()

		if strings.TrimSpace(name) == "" {
			logrus.Fatal("invalid migration name")
		}

		m := jen.NewFile("migrations")

		m.Comment("Auto-generated file: https://github.com/semirm-dev/gomigrate")
		m.Comment("Feel free to edit")
		m.Line()

		m.Comment(name + " migration")
		m.Type().Id(name).Struct()

		m.Comment("Apply migration")
		m.Func().Params(jen.Id("mig").Id("*" + name)).Id("Apply").Params(jen.Id("db").Op("*").Qual("github.com/jinzhu/gorm", "DB")).Block(
			jen.Qual("github.com/sirupsen/logrus", "Info").Call(jen.Lit("Applying migration")),
		)

		m.Comment("Rollback migration")
		m.Func().Params(jen.Id("mig").Id("*" + name)).Id("Rollback").Params(jen.Id("db").Op("*").Qual("github.com/jinzhu/gorm", "DB")).Block(
			jen.Qual("github.com/sirupsen/logrus", "Info").Call(jen.Lit("Rolling back migration")),
		)

		ts := time.Now().Unix()

		m.Comment("Timestamp when migration was created")
		m.Func().Params(jen.Id("mig").Id("*" + name)).Id("Timestamp").Params().Int64().Block(
			jen.Return(jen.Lit(ts)),
		)

		m.Comment("Name of migration")
		m.Func().Params(jen.Id("mig").Id("*" + name)).Id("Name").Params().String().Block(
			jen.Return(jen.Lit(name)),
		)

		n := strings.ToLower(fmt.Sprint(ts) + "_" + name + ".go")
		c := []byte(fmt.Sprintf("%#v", m))

		writeFileContent(migrationsDest+"/"+n, c)

		logrus.Info("migration created")
	},
}
