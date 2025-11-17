package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type JSONStorage struct {
	FilePath string
}

func NewJSONStorage(filePath string) *JSONStorage {
	return &JSONStorage{FilePath: filePath}
}

func (s *JSONStorage) EnsureFile() error {
	dir := filepath.Dir(s.FilePath)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	_, err = os.Stat(s.FilePath)
	if os.IsNotExist(err) {
		return os.WriteFile(s.FilePath, []byte("[]"), 0644)
	}

	return nil
}

func (s *JSONStorage) Load() ([]byte, error) {
	if err := s.EnsureFile(); err != nil {
		return nil, err
	}

	return os.ReadFile(s.FilePath)
}

func (s *JSONStorage) Save(data []byte) error {
	if len(data) == 0 {
		return errors.New("data cannot be empty")
	}

	return os.WriteFile(s.FilePath, data, 0644)
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
