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
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(info.Name(), r.extension) {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("couldn't read file %s. %w", info.Name(), err)
		}
		parser(path, data)
		return nil
	})
	if err != nil {
		return fmt.Errorf("couldn't read directory %s. %w", path, err)
	}
	return nil
}
