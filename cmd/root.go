package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "m8",
	Short: "Simple SQL migration tool for PostgreSQL",
	Long: `m8 is a lightweight CLI tool for managing PostgreSQL schema migrations.
It helps you create versioned SQL migration files and apply them in order.
Each migration consists of an "up" script to apply changes and a "down" script to roll them back.`,
}

func init() {}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
