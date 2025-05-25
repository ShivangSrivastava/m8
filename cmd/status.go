package cmd

import (
	"github.com/ShivangSrivastava/m8/internal/cli"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

// Defines the `status` CLI command to show migration state.
// Binds to `cli.Status` to run logic and print applied/pending migrations.
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of all applied and pending migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.Status(cmd, args)
	},
}

// Registers the command to the root so it's available in CLI.
func init() {
	rootCmd.AddCommand(statusCmd)
}
