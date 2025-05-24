package fs_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ShivangSrivastava/m8/internal/infra/fs"
)

// TestCreateMigrations verifies that CreateMigrations correctly generates the expected
// .up.sql and .down.sql files in the given directory. It ensures that file naming follows
// the version_name pattern and that files are physically created on disk. This test guards
// against regressions in file output logic and naming conventions.
func TestCreateMigrations(t *testing.T) {
	dir := t.TempDir()

	creator := &fs.FileCreator{dir}

	version := "20250524151203"
	name := "test_migrations"

	if err := creator.CreateMigrations(version, name); err != nil {
		t.Fatalf("CreateMigrations failed: %v", err)
	}

	upPath := filepath.Join(dir, version+"_"+name+".up.sql")
	downPath := filepath.Join(dir, version+"_"+name+".down.sql")

	if _, err := os.Stat(upPath); os.IsNotExist(err) {
		t.Errorf("Expected file %x not found", upPath)
	}

	if _, err := os.Stat(downPath); os.IsNotExist(err) {
		t.Errorf("Expected file %x not found", downPath)
	}
}
