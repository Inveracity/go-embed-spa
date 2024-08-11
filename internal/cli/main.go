package cli

import (
	"github.com/inveracity/go-embed-spa/internal/client"
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
)

type App struct {
	Version string
	Client  client.Client
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
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			InitializeEnv(cmd)
			endpoint, err := cmd.Flags().GetString("endpoint")
			if err != nil {
				return err
			}
			a.Client = *client.New(endpoint)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	rootCmd.PersistentFlags().String("endpoint", "http://localhost:3000", "--endpoint=http://localhost:3000")

	rootCmd.AddCommand((&CmdVersion{}).Command())
	rootCmd.AddCommand((&CmdHello{Client: &a.Client}).Command())
	rootCmd.AddCommand((&CmdServer{}).Command())
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true

	return rootCmd.Execute()
}
