package cli

import (
	"github.com/spf13/cobra"

	"github.com/inveracity/go-embed-spa/internal/client"
)

type CmdMemory struct {
	Client *client.Client
}

func (h *CmdMemory) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "memory",
		Short: "stream memory data",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := h.Client.Memory()
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().SortFlags = false
	return cmd
}
