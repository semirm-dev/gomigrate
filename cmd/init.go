package cmd

func init() {
	parseConfigFile(cmdDest + "/config.yml")

	Migration.AddCommand(Create)
	Migration.AddCommand(Template)
}
