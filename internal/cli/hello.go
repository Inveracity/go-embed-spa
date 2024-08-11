package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/inveracity/go-embed-spa/internal/client"
)

type CmdHello struct {
	Client *client.Client
}

func (h *CmdHello) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "Responds with \"world\"",
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := h.Client.Hello()
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().SortFlags = false
	return cmd
}
