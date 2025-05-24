package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileCreater struct {
	Dir string
}

func (f *FileCreater) CreateMigrations(version, name string) error {
	migrationName := fmt.Sprintf("%s_%s", version, name)
	upFile, downFile := migrationName+".up.sql", migrationName+".down.sql"

	if _, err := os.Create(filepath.Join(f.Dir, upFile)); err != nil {
		return err
	}

	if _, err := os.Create(filepath.Join(f.Dir, downFile)); err != nil {
		return err
	}
	return nil
}
