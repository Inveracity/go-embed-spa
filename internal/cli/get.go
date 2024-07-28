package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Get struct {
}

func (g *Get) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get data from server",

		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := g.getData()
			fmt.Println(result)
			return err
		},
	}

	cmd.Flags().SortFlags = false
	return cmd
}

func (g *Get) getData() (string, error) {
	return "data", nil
}
