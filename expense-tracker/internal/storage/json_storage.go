package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type JSONStorage struct {
	filepath string
}

func NewJSONStorage(filepath string) *JSONStorage {
	return &JSONStorage{filepath: filepath}
}

func (s *JSONStorage) EnsureFile() error {
	dir := filepath.Dir(s.filepath)

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	_, err := os.Stat(s.filepath)
	if os.IsNotExist(err) {
		return os.WriteFile(s.filepath, []byte("[]"), 0o644)
	}
	return nil
}

func (s *JSONStorage) Load() ([]byte, error) {
	if err := s.EnsureFile(); err != nil {
		return nil, err
	}

	return os.ReadFile(s.filepath)
}

func (s *JSONStorage) Save(data []byte) error {
	if len(data) == 0 {
		return errors.New("data cannot be empty")
	}

	dir := filepath.Dir(s.filepath)
	tmpFile, err := os.CreateTemp(dir, "temp-*.json")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(data); err != nil {
		return err
	}
	if err := tmpFile.Chmod(0o644); err != nil {
		return err
	}
	if err := tmpFile.Sync(); err != nil {
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return err
	}
	return os.Rename(tmpFile.Name(), s.filepath)
}

func (s *JSONStorage) LoadJSON(v any) error {
	data, err := s.Load()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (s *JSONStorage) SaveJSON(v any) error {
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	return s.Save(out)
}
