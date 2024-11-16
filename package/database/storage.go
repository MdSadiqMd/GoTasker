package database

import (
	"encoding/json"
	"os"

	"github.com/MdSadiqMd/GoTasker/package/types"
)

// Wrapper of Database type from original types
type Storage[T any] struct {
	types.Storage[T]
}

func NewStorage[T any](fileName string) *types.Storage[T] {
	return &types.Storage[T]{FileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "")
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
