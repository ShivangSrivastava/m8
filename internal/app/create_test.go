package app_test

import (
	"errors"
	"testing"

	"github.com/ShivangSrivastava/m8/internal/app"
)

type mockCreator struct {
	called  bool
	version string
	name    string
	fail    bool
}

func (m *mockCreator) CreateMigrations(version, name string) error {
	m.called = true
	m.version = version
	m.name = name
	if m.fail {
		return errors.New("mock failure")
	}
	return nil
}

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

