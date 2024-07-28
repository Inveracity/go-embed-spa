package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionNumber = "dev"

type Version struct{}

func (c *Version) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(versionNumber)
		},
	}

	return cmd
}
