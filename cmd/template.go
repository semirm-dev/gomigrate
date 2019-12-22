package cmd

import (
	"fmt"
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

		var pkg = cmd.Flag("pkg").Value.String()

		if strings.TrimSpace(pkg) == "" {
			logrus.Fatal("invalid github package url")
		}

		name := "cmd"
		m := jen.NewFile(name)

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
					jen.Qual("github.com/"+pkg+"/migrations", "Run").Call(),
					jen.Qual("github.com/sirupsen/logrus", "Info").Call(jen.Lit("migrations applied")),
				),
			},
		)

		c := []byte(fmt.Sprintf("%#v", m))

		writeFileContent(cmdDest+"/migration.go", c)

		logrus.Info("templates generated")
	},
}
