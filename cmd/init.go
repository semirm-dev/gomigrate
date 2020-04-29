package cmd

func init() {
	Migration.AddCommand(Create)
	Migration.AddCommand(Template)
}
