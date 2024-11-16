package database

import (
	"encoding/json"
	"os"

	"github.com/MdSadiqMd/GoTasker/package/types"
)

// Wrapper around the original Storage type
type Storage[T any] struct {
	types.Storage[T]
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{Storage: types.Storage[T]{FileName: fileName}}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
