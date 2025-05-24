package fs_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ShivangSrivastava/m8/internal/infra/fs"
)

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
