package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionNumber = "dev"

type CmdVersion struct{}

func (c *CmdVersion) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(versionNumber)
		},
	}

	return cmd
}
