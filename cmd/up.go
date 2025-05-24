package cmd

import (
	"github.com/ShivangSrivastava/m8/internal/cli"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

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
