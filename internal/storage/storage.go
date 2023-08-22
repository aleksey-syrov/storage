package storage

import (
	"fmt"
	"github.com/aleksey-syrov/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{Files: make(map[uuid.UUID]*file.File)}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	uploadedFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.Files[uploadedFile.ID] = uploadedFile

	return uploadedFile, err
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	fmt.Println("here", s.Files)
	if foundFile, ok := s.Files[fileID]; ok {
		return foundFile, nil
	}

	return &file.File{}, fmt.Errorf("file %v not found", fileID)
}
