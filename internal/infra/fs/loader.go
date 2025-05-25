package fs

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ShivangSrivastava/m8/internal/core"
)

type Direction string

const (
	Up   Direction = "up"
	Down Direction = "down"
)

type FileLoader struct {
	Dir       string
	Direction Direction
}

// LoadMigrations reads all migration files from the filesystem in the specified direction (up/down).
// It assumes migration files follow a naming pattern like `<version>_<name>.up.sql` or `.down.sql`.
// This function sorts the files to ensure they are applied in the correct order,
// extracts the version from the filename, and returns a list of structured migrations.
func (f *FileLoader) LoadMigrations() ([]core.Migration, error) {
	pattern := "*.up.sql"
	if f.Direction == Down {
		pattern = "*.down.sql"
	}

	files, err := filepath.Glob(filepath.Join(f.Dir, pattern))
	if err != nil {
		return nil, err
	}

	sort.Strings(files)

	var migrations []core.Migration

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		version := strings.Split(filepath.Base(file), "_")[0]

		if f.Direction == Up {
			migrations = append(migrations, core.Migration{
				Name:    filepath.Base(file),
				Version: version,
				UpSQL:   string(content),
			})
		} else {
			migrations = append(migrations, core.Migration{
				Name:    filepath.Base(file),
				Version: version,
				DownSQL: string(content),
			})
		}
	}

	return migrations, nil
}

// LoadMigration loads a specific down migration file by version.
// Reads the file content and returns a Migration with DownSQL.
// Used during revert to get the rollback SQL for the latest version.
func (f *FileLoader) LoadMigration(version string) (core.Migration, error) {
	pattern := version + "*.down.sql"
	files, err := filepath.Glob(filepath.Join(f.Dir, pattern))
	if err != nil {
		return core.Migration{}, err
	}

	file := files[0]
	content, err := os.ReadFile(file)
	if err != nil {
		return core.Migration{}, err
	}
	result := core.Migration{
		Name:    filepath.Base(file),
		Version: version,
		DownSQL: string(content),
	}
	return result, nil
}
