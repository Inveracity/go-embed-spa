package cli

import (
	"github.com/inveracity/go-embed-spa/internal/server"
	"github.com/spf13/cobra"
)

type Server struct {
}

func (s *Server) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "run server",

		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetUint("port")
			return s.run(port)

		},
	}

	cmd.Flags().SortFlags = false
	cmd.PersistentFlags().UintP("port", "p", 3000, "--port=3000 | -p 3000")
	return cmd
}

func (s *Server) run(port uint) error {

	server.Server(port)

	return nil
}
