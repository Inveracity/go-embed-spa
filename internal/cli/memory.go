package cli

import (
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"

	"github.com/inveracity/go-embed-spa/internal/client"
)

type CmdMemory struct {
	Client *client.Client
}

func (h *CmdMemory) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "memory",
		Short: "Stream memory data",
		Long: dedent.Dedent(`
			Example:
			  cli memory | jq .[].usedPercent`,
		),
		SilenceErrors: true, // Bubble up errors instead
		SilenceUsage:  true,
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
