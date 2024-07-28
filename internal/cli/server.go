package cli

import (
	"fmt"
	"net/http"

	"github.com/inveracity/go-embed-spa/internal/server"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

type Server struct {
}

func (s *Server) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "run server",

		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetUint("port")
			apiport, _ := cmd.Flags().GetUint("apiport")
			return s.run(port, apiport)

		},
	}

	cmd.Flags().SortFlags = false
	cmd.PersistentFlags().UintP("port", "p", 3000, "--port=3000 | -p 3000")
	cmd.PersistentFlags().Uint("apiport", 3001, "--apiport=3001")
	return cmd
}

func (s *Server) run(port uint, apiport uint) error {

	server01 := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", port),
		Handler: server.Server(),
	}

	server02 := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", apiport),
		Handler: server.ServerAPI(),
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
