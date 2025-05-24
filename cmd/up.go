package cmd

import (
	"github.com/ShivangSrivastava/m8/internal/cli"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

// upCmd defines the CLI command that triggers applying all pending "up" migrations to the database.
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all pending database migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.Up(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
