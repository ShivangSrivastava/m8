package cmd

import (
	"github.com/ShivangSrivastava/m8/internal/cli"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

// This command is part of the migration tool’s rollback capability,
// allowing users to safely undo the latest schema change.
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of all applied and pending migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.Status(cmd, args)
	},
}

// Registering downCmd in init ensures it becomes part of the CLI interface
// without requiring manual invocation elsewhere.
func init() {
	rootCmd.AddCommand(statusCmd)
}
