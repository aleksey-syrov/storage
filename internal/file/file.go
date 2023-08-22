package file

import "github.com/google/uuid"

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		Name: filename,
		Data: blob,
		ID:   fileId,
	}, nil
}
