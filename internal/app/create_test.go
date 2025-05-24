package app_test

import (
	"errors"
	"testing"

	"github.com/ShivangSrivastava/m8/internal/app"
)

// mockCreator is a test double implementing MigrationCreator for verifying
// CreateService behavior without touching the filesystem. It records call parameters
// and can simulate failures.
type mockCreator struct {
	called  bool
	version string
	name    string
	fail    bool
}

// CreateMigrations records the invocation parameters and optionally returns an error
// to test error handling in CreateService.
func (m *mockCreator) CreateMigrations(version, name string) error {
	m.called = true
	m.version = version
	m.name = name
	if m.fail {
		return errors.New("mock failure")
	}
	return nil
}

// TestCreateFile_Success ensures CreateService correctly sanitizes the migration name,
// generates a version, and calls the Creator without error.
// It verifies integration of name formatting and delegation logic.
func TestCreateFile_Success(t *testing.T) {
	mock := &mockCreator{}
	service := &app.CreateService{
		Name:    "my.one-feature 123",
		Creator: mock,
	}

	err := service.CreateFile()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !mock.called {
		t.Fatal("Expected CreateMigrations to be called")
	}
	if mock.name != "my_one_feature_123" {
		t.Errorf("Expected sanitized name to be 'my_one_feature_123', got '%s'", mock.name)
	}
}
