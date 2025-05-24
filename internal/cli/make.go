package cli

import (
	"log"

	"github.com/ShivangSrivastava/m8/internal/app"
	"github.com/ShivangSrivastava/m8/internal/infra/fs"
	"github.com/spf13/cobra"
)

func Make(cmd *cobra.Command, args []string) error {
	creator := &fs.FileCreater{Dir: "migrations"}
	service := &app.CreateService{
		Name:    args[0],
		Creator: creator,
	}
	if err := service.CreateFile(); err != nil {
		log.Fatalln(err)
	}
	return nil
}
