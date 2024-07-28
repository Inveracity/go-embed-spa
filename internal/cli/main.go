package cli

import (
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
)

type App struct {
	Version string
}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "cli interfaces with server-api",
		Long: dedent.Dedent(`
			usage: cli --help
			`,
		),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand((&Version{}).Command())
	rootCmd.AddCommand((&Get{}).Command())
	rootCmd.AddCommand((&Server{}).Command())

	return rootCmd.Execute()
}
