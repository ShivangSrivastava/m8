package cmd

import (
	"github.com/ShivangSrivastava/m8/internal/cli"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Revert the last applied migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.Down(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
