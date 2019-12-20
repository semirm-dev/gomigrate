package cmd

import (
	"fmt"

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

		if name == "" {
			logrus.Fatal("invalid migration name")
		}

		m := jen.NewFile("migrations")

		m.Type().Id(name).Struct()

		m.Func().Params(jen.Id("mig").Id("*" + name)).Id("Apply").Params().Block(
			jen.Qual("github.com/sirupsen/logrus", "Info").Call(jen.Lit("Applying migration")),
		)

		m.Func().Params(jen.Id("mig").Id("*" + name)).Id("Rollback").Params().Block(
			jen.Qual("github.com/sirupsen/logrus", "Info").Call(jen.Lit("Rollback migration")),
		)

		fmt.Printf("%#v", m)

		logrus.Info("migration created")
	},
}
