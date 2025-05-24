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
