package cli

import (
	"github.com/inveracity/go-embed-spa/internal/client"
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
)

type App struct {
	Client  client.Client
	NoColor bool
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
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			InitializeEnv(cmd)
			a.NoColor, err = cmd.Flags().GetBool("no-color")
			if err != nil {
				return err
			}
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
	rootCmd.PersistentFlags().String("endpoint", "http://localhost:3000", "set the http endpoint of the server")
	rootCmd.PersistentFlags().Bool("no-color", false, "disable colored output")

	rootCmd.AddCommand((&CmdVersion{}).Command())
	rootCmd.AddCommand((&CmdHello{Client: &a.Client}).Command())
	rootCmd.AddCommand((&CmdMemory{Client: &a.Client}).Command())
	rootCmd.AddCommand((&CmdServer{}).Command())
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true

	return rootCmd.Execute()
}
