package cmd

import (
	"github.com/ShivangSrivastava/m8/internal/cli"
	"github.com/spf13/cobra"
)

// makeCmd defines the CLI command for generating new SQL migration files.
// It ensures a migration name is provided and delegates execution to the cli.Make handler.
// This command integrates with the root command to expose migration creation functionality.
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
