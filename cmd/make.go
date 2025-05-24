package cmd

import (
	"github.com/ShivangSrivastava/m8/internal/cli"
	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make [name]",
	Short: "Create new `.up.sql` and `.down.sql` files",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.Make(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}
