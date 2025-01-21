package writer

import (
	"encoding/json"
	"fmt"
	"os"
)

type Writer interface {
	WriteFile(obj interface{}, filePath string, enableIndentation bool) error
}

type writer struct{}

func New() Writer {
	return &writer{}
}

func (w *writer) WriteFile(obj interface{}, filePath string, enableIndentation bool) error {
	var (
		err  error
		data []byte
	)
	if enableIndentation {
		data, err = json.MarshalIndent(obj, "", "  ")
	} else {
		data, err = json.Marshal(obj)
	}
	if err != nil {
		return fmt.Errorf("couldn't marshal output %v", err)
	}

	err = os.WriteFile(filePath, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("couldn't write file %s. %v", filePath, err)
	}
	return nil
}
