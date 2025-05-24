package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileCreator stores dir name, where the pair of migration files i.e. *.up.sql and *.down.sql will be created
type FileCreator struct {
	Dir string
}

// CreateMigrations generates empty .up.sql and .down.sql files for a given migration.
// It standardizes naming by combining version and name, and ensures the files are placed
// in the configured directory. This helps enforce consistency across migrations
// and avoids manual file creation errors.
func (f *FileCreator) CreateMigrations(version, name string) error {
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
