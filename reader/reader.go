package reader

import (
	"fmt"
	"os"
	"path/filepath"
)

// Reader reads boards from files
type Reader interface {
	// ReadDirectory reads all files in a directory and calls the parser on the read data
	ReadDirectory(path string, parser func(filePath string, data []byte)) error
}

func New() Reader {
	return &reader{}
}

type reader struct{}

func (r *reader) ReadDirectory(path string, parser func(filePath string, data []byte)) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("couldn't read directory %s. %v", path, err)
	}

	for _, e := range entries {
		filePath := filepath.Join(path, e.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("couldn't read file %s. %v", e.Name(), err)
		}
		parser(filePath, data)
	}

	return nil
}
