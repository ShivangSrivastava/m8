package fs_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ShivangSrivastava/m8/internal/infra/fs"
)

// TestLoadMigrations verifies that FileLoader correctly loads both up and down migration files from disk.
// It creates mock migration files in a temporary directory to simulate real input,
// ensuring that the loader reads the correct direction based on configuration,
// parses filenames to extract the correct version,
// and returns non-empty, valid migration entries.
func TestLoadMigrations(t *testing.T) {
	dir := t.TempDir()
	version := "20250524160302"
	name := "test_loader"
	if _, err := os.Create(filepath.Join(dir, version+"_"+name+".up.sql")); err != nil {
		t.Fatalf("TestLoadMigrations failed: %v", err)
	}

	if _, err := os.Create(filepath.Join(dir, version+"_"+name+".down.sql")); err != nil {
		t.Fatalf("TestLoadMigrations failed: %v", err)
	}
	upLoader := &fs.FileLoader{dir, fs.Up}
	downLoader := &fs.FileLoader{dir, fs.Down}

	upMigrFiles, err := upLoader.LoadMigrations()
	if err != nil {
		t.Fatalf("TestLoadMigrations failed: %v", err)
	}

	downMigrFiles, err := downLoader.LoadMigrations()
	if err != nil {
		t.Fatalf("TestLoadMigrations failed: %v", err)
	}

	if len(upMigrFiles) == 0 {
		t.Fatalf("Expected at least one up migration file, found none")
	}
	if len(downMigrFiles) == 0 {
		t.Fatalf("Expected at least one down migration file, found none")
	}

	if upMigrFiles[0].Version != version {
		t.Errorf("Expected: %s, Found: %s", version, upMigrFiles[0].Version)
	}
	if downMigrFiles[0].Version != version {
		t.Errorf("Expected: %s, Found: %s", version, downMigrFiles[0].Version)
	}
}

// Creates a temporary down migration file to simulate a real revert case.
// Verifies LoadMigration reads correct version, name, and SQL content.
// Ensures DownSQL matches expected value from the test file.
func TestLoadMigration(t *testing.T) {
	dir := t.TempDir()
	version := "20250524160302"
	name := "test_loader"
	filename := version + "_" + name + ".down.sql"
	filePath := filepath.Join(dir, filename)
	content := "DROP TABLE cars;"

	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		t.Fatalf("Failed to write migration file: %v", err)
	}

	loader := &fs.FileLoader{Dir: dir}

	migration, err := loader.LoadMigration(version)
	if err != nil {
		t.Fatalf("LoadMigration returned error: %v", err)
	}

	if migration.Version != version {
		t.Errorf("Expected version %s, got %s", version, migration.Version)
	}
	if migration.Name != filename {
		t.Errorf("Expected name %s, got %s", filename, migration.Name)
	}
	if migration.DownSQL != content {
		t.Errorf("Expected content %q, got %q", content, migration.DownSQL)
	}
}
