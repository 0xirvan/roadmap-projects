package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestEnsureFile tests the EnsureFile method for various scenarios.
func TestEnsureFile(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(t *testing.T, path string)
		wantErr   bool
		checkFile func(t *testing.T, path string)
	}{
		{
			name: "creates directory and file",
			setup: func(t *testing.T, path string) {
				// no setup needed
			},
			wantErr: false,
			checkFile: func(t *testing.T, path string) {
				if _, err := os.Stat(path); os.IsNotExist(err) {
					t.Error("file was not created")
				}
				data, err := os.ReadFile(path)
				if err != nil {
					t.Fatalf("failed to read file: %v", err)
				}
				if string(data) != "[]" {
					t.Errorf("file content = %s, want []", string(data))
				}
			},
		},
		{
			name: "succeeds when file already exists",
			setup: func(t *testing.T, path string) {
				if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
					t.Fatalf("failed to create directory: %v", err)
				}
				if err := os.WriteFile(path, []byte(`[{"id":1}]`), 0o644); err != nil {
					t.Fatalf("failed to write file: %v", err)
				}
			},
			wantErr: false,
			checkFile: func(t *testing.T, path string) {
				data, err := os.ReadFile(path)
				if err != nil {
					t.Fatalf("failed to read file: %v", err)
				}
				if string(data) != `[{"id":1}]` {
					t.Error("existing file was modified")
				}
			},
		},
		{
			name: "succeeds when called multiple times",
			setup: func(t *testing.T, path string) {
				// first call will create it
				storage := NewJSONStorage(path)
				if err := storage.EnsureFile(); err != nil {
					t.Fatalf("first EnsureFile failed: %v", err)
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "subdir", "test.json")

			if tt.setup != nil {
				tt.setup(t, path)
			}

			storage := NewJSONStorage(path)
			err := storage.EnsureFile()

			if (err != nil) != tt.wantErr {
				t.Errorf("EnsureFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.checkFile != nil {
				tt.checkFile(t, path)
			}
		})
	}
}

// TestLoad tests the Load method for various scenarios.
func TestLoad(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(t *testing.T, path string)
		wantErr     bool
		wantContent string
	}{
		{
			name: "successfully reads existing file",
			setup: func(t *testing.T, path string) {
				if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
					t.Fatalf("failed to create directory: %v", err)
				}
				if err := os.WriteFile(path, []byte(`[{"id":1}]`), 0o644); err != nil {
					t.Fatalf("failed to write file: %v", err)
				}
			},
			wantErr:     false,
			wantContent: `[{"id":1}]`,
		},
		{
			name:        "creates file if not exists",
			setup:       func(t *testing.T, path string) {},
			wantErr:     false,
			wantContent: "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "test.json")

			if tt.setup != nil {
				tt.setup(t, path)
			}

			storage := NewJSONStorage(path)
			got, err := storage.Load()

			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if string(got) != tt.wantContent {
				t.Errorf("Load() = %s, want %s", string(got), tt.wantContent)
			}
		})
	}
}

// TestSave tests the Save method for various scenarios.
func TestSave(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		wantErr  bool
		wantData string
	}{
		{
			name:     "successfully saves data",
			data:     []byte(`[{"id":1,"name":"test"}]`),
			wantErr:  false,
			wantData: `[{"id":1,"name":"test"}]`,
		},
		{
			name:     "returns error on empty data",
			data:     []byte{},
			wantErr:  true,
			wantData: "",
		},
		{
			name:     "overwrites existing data",
			data:     []byte(`[{"id":2},{"id":3}]`),
			wantErr:  false,
			wantData: `[{"id":2},{"id":3}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "test.json")
			storage := NewJSONStorage(path)

			err := storage.Save(tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.wantData != "" {
				got, err := os.ReadFile(path)
				if err != nil {
					t.Fatalf("failed to read saved file: %v", err)
				}
				if string(got) != tt.wantData {
					t.Errorf("saved data = %s, want %s", string(got), tt.wantData)
				}
			}
		})
	}
}

// TestLoadJSON tests the LoadJSON method for JSON unmarshaling.
func TestLoadJSON(t *testing.T) {
	type testData struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	tests := []struct {
		name    string
		setup   func(t *testing.T, path string)
		wantLen int
		wantErr bool
	}{
		{
			name: "successfully unmarshals JSON",
			setup: func(t *testing.T, path string) {
				if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
					t.Fatalf("failed to create directory: %v", err)
				}
				data := []testData{{ID: 1, Name: "test"}, {ID: 2, Name: "test2"}}
				jsonData, err := json.MarshalIndent(data, "", "  ")
				if err != nil {
					t.Fatalf("failed to marshal JSON: %v", err)
				}
				if err := os.WriteFile(path, jsonData, 0o644); err != nil {
					t.Fatalf("failed to write file: %v", err)
				}
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "creates empty array if file not exists",
			setup:   func(t *testing.T, path string) {},
			wantLen: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "test.json")

			if tt.setup != nil {
				tt.setup(t, path)
			}

			storage := NewJSONStorage(path)
			var got []testData
			err := storage.LoadJSON(&got)

			if (err != nil) != tt.wantErr {
				t.Errorf("LoadJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != tt.wantLen {
				t.Errorf("LoadJSON() length = %d, want %d", len(got), tt.wantLen)
			}
		})
	}
}

// TestSaveJSON tests the SaveJSON method for JSON marshaling.
func TestSaveJSON(t *testing.T) {
	type testData struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	tests := []struct {
		name    string
		data    any
		wantErr bool
		verify  func(t *testing.T, path string)
	}{
		{
			name: "successfully marshals and saves JSON",
			data: []testData{{ID: 1, Name: "test"}, {ID: 2, Name: "test2"}},
			verify: func(t *testing.T, path string) {
				storage := NewJSONStorage(path)
				var loaded []testData
				if err := storage.LoadJSON(&loaded); err != nil {
					t.Fatalf("failed to load JSON: %v", err)
				}
				if len(loaded) != 2 {
					t.Errorf("loaded data length = %d, want 2", len(loaded))
				}
				if loaded[0].ID != 1 || loaded[1].ID != 2 {
					t.Error("loaded data does not match saved data")
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "test.json")
			storage := NewJSONStorage(path)

			err := storage.SaveJSON(tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("SaveJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.verify != nil {
				tt.verify(t, path)
			}
		})
	}
}
