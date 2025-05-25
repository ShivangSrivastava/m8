package cmd

import (
	"github.com/ShivangSrivastava/m8/internal/cli"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

// This command is part of the migration tool’s rollback capability,
// allowing users to safely undo the latest schema change.
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Revert the last applied migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.Down(cmd, args)
	},
}

// Registering downCmd in init ensures it becomes part of the CLI interface
// without requiring manual invocation elsewhere.
func init() {
	rootCmd.AddCommand(downCmd)
}
