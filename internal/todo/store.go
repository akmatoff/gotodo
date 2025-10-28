package todo

import (
	"encoding/json"
	"os"
)

type Store interface {
	Save(todos []Todo) error
	Load() ([]Todo, error)
}

type FileStore struct {
	FileName string
}

func (fs *FileStore) Load() ([]Todo, error) {
	data, err := os.ReadFile(fs.FileName)

	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}

		return nil, err
	}

	var todos []Todo

	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (fs *FileStore) Save(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(fs.FileName, data, 0644)
}
