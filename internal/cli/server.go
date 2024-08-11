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
		Short:  "run server",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetUint("server-port")
			return s.run(port)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.PersistentFlags().UintP("server-port", "p", 3000, "--server-port=3000 | -p 3000")
	return cmd
}

func (s *CmdServer) run(port uint) error {
	server.Server(port)
	return nil
}
