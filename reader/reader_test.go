package reader

import (
	"path/filepath"
	"testing"
)

func Test_reader_ReadDirectory(t *testing.T) {
	t.Run("Should read all files in a valid directory", func(t *testing.T) {
		cnt := 0
		err := New(".json").ReadDirectory(filepath.Join("..", "testdata", "multiple_files"), func(filePath string, data []byte) {
			cnt++
		})
		if err != nil {
			t.Errorf("ReadDirectory returned unexpected error: %v", err)
		}
		if cnt != 2 {
			t.Errorf("ReadDirectory returned %d files, expected 2", cnt)
		}
	})
	t.Run("Should return an error if the directory does not exist", func(t *testing.T) {
		err := New(".json").ReadDirectory(filepath.Join("..", "invalid"), func(filePath string, data []byte) {})
		if err == nil {
			t.Errorf("ReadDirectory returned nil error")
		}
	})
	t.Run("Should ignore non-json files", func(t *testing.T) {
		cnt := 0
		err := New(".json").ReadDirectory(filepath.Join("..", "testdata", "nonjson"), func(filePath string, data []byte) {
			cnt++
		})
		if err != nil {
			t.Errorf("ReadDirectory returned unexpected error: %v", err)
		}
		if cnt != 0 {
			t.Errorf("ReadDirectory returned %d files, expected 2", cnt)
		}
	})
}
