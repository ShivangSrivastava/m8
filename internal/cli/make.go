package cli

import (
	"log"

	"github.com/ShivangSrivastava/m8/internal/app"
	"github.com/ShivangSrivastava/m8/internal/infra/fs"
	"github.com/spf13/cobra"
)

// Make is the CLI handler that creates migration files based on user input.
// It sets up the file system creator and the service, then delegates file creation.
// Exits the program on error to provide immediate feedback in the CLI.
func Make(cmd *cobra.Command, args []string) error {
	creator := &fs.FileCreator{Dir: "migrations"}
	service := &app.CreateService{
		Name:    args[0],
		Creator: creator,
	}
	if err := service.CreateFile(); err != nil {
		log.Fatalln(err)
	}
	return nil
}
