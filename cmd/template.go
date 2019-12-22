package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Template migration command
var Template = &cobra.Command{
	Use:   "tpl",
	Short: "Generate templates",
	Long:  `Generate templates`,
	Run: func(cmd *cobra.Command, args []string) {
		createPathIfNotExists(cmdDest)

		createMigrationInterface()

		createRegisterMigrationsCollection()

		createApplyCmd(cmd)

		logrus.Info("templates generated")
	},
}

func createMigrationInterface() {
	from := "https://raw.githubusercontent.com/semirm-dev/gomigrate/master/cmd/migration.tpl"
	dest := migrationsDest + "/migration.go"

	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := downloadTpl(from, dest); err != nil {
			logrus.Fatal("failed to get migration.tpl: ", err)
		}
	}
}

func createRegisterMigrationsCollection() {
	dest := migrationsDest + "/registermigrations.go"

	if _, err := os.Stat(dest); os.IsNotExist(err) {
		m := jen.NewFile(migrationsDest)

		m.Comment("Auto-generated file: https://github.com/semirm-dev/gomigrate")
		m.Comment("Feel free to edit")
		m.Line()

		m.Comment("Collection with all migrations")
		m.Var().Id("Collection").Op("=").Index().Qual("", "Migration").Block()

		c := []byte(fmt.Sprintf("%#v", m))

		writeFileContent(dest, c)
	}
}

func createApplyCmd(cmd *cobra.Command) {
	dest := cmdDest + "/migration.go"

	if _, err := os.Stat(dest); os.IsNotExist(err) {
		var pkg = cmd.Flag("pkg").Value.String()

		if strings.TrimSpace(pkg) == "" {
			logrus.Fatal("invalid github package url")
		}

		m := jen.NewFile(cmdDest)

		m.Comment("Auto-generated file: https://github.com/semirm-dev/gomigrate")
		m.Comment("Feel free to edit")
		m.Line()

		m.Func().Id("init").Params().Block(
			jen.Qual("", "Migration.AddCommand").Call(jen.Id("Apply")),
		)

		m.Comment("Migration command")
		m.Var().Id("Migration").Op("= &").Qual("github.com/spf13/cobra", "Command").Values(
			jen.Dict{
				jen.Id("Use"):   jen.Lit(""),
				jen.Id("Short"): jen.Lit("Migrations tool"),
				jen.Id("Long"):  jen.Lit("`Migrations tool`"),
				jen.Id("Run"):   jen.Func().Params(jen.Id("cmd").Op("*").Qual("github.com/spf13/cobra", "Command"), jen.Id("agrs").Index().String()).Block(),
			},
		)

		m.Comment("Apply command")
		m.Var().Id("Apply").Op("= &").Qual("github.com/spf13/cobra", "Command").Values(
			jen.Dict{
				jen.Id("Use"):   jen.Lit("migrate"),
				jen.Id("Short"): jen.Lit("Apply migrations"),
				jen.Id("Long"):  jen.Lit("`Apply migrations`"),
				jen.Id("Run"): jen.Func().Params(jen.Id("cmd").Op("*").Qual("github.com/spf13/cobra", "Command"), jen.Id("agrs").Index().String()).Block(
					jen.Qual("github.com/"+pkg+"/"+migrationsDest, "Run").Call(),
					jen.Qual("github.com/sirupsen/logrus", "Info").Call(jen.Lit("migrations applied")),
				),
			},
		)

		c := []byte(fmt.Sprintf("%#v", m))

		writeFileContent(dest, c)
	}
}
