package cli

import (
	"github.com/inveracity/go-embed-spa/internal/server"
	"github.com/spf13/cobra"
)

type CmdServer struct {
}

func (s *CmdServer) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "server",
		Short:  "Run server",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetUint("server-port")
			bind, _ := cmd.Flags().GetString("server-addr")
			server.Server(bind, port)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.PersistentFlags().UintP("server-port", "p", 3000, "set the listening port of the server")
	cmd.PersistentFlags().StringP("server-addr", "a", "0.0.0.0", "set the bind address of the server")
	return cmd
}
