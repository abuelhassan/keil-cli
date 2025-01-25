package reader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Reader interface {
	// ReadDirectory reads all files in a directory and calls the parser on the read data
	ReadDirectory(path string, parser func(filePath string, data []byte)) error
}

func New(extension string) Reader {
	return &reader{extension: extension}
}

type reader struct {
	extension string
}

func (r *reader) ReadDirectory(path string, parser func(filePath string, data []byte)) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("couldn't read directory %s. %w", path, err)
	}

	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), r.extension) {
			continue
		}
		filePath := filepath.Join(path, e.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("couldn't read file %s. %w", e.Name(), err)
		}
		parser(filePath, data)
	}

	return nil
}
