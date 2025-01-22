package main

import (
	"github.com/abuelhassan/keil-cli/board"
	"reflect"
	"testing"
)

func Test_mergeAction(t *testing.T) {
	type args struct {
		dir               string
		enableIndentation bool
		outputFile        string
		rdr               func() *mockReader
		wrt               func() *mockWriter
	}
	tests := []struct {
		name    string
		args    args
		want    board.Summary
		wantErr bool
	}{
		{
			name: "Should read, merge and write summary",
			args: args{
				dir: "testdata",
				rdr: func() *mockReader {
					return &mockReader{
						ReadDirectoryFn: func(path string, parser func(filePath string, data []byte)) error {
							parser("file.json", []byte(`{"boards":[{"name":"name","vendor":"vendor"}]}`))
							return nil
						},
					}
				},
				wrt: func() *mockWriter {
					m := &mockWriter{}
					m.WriteFileFn = func(obj interface{}, filePath string, enableIndentation bool) error {
						m.WrittenSummary = obj.(board.Summary)
						return nil
					}
					return m
				},
			},
			want: board.Summary{
				Boards: []board.Board{
					{
						Name:   "name",
						Vendor: "vendor",
					},
				},
				Metadata: board.Metadata{
					TotalVendors: 1,
					TotalBoards:  1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWrt := tt.args.wrt()
			mergeAction(tt.args.dir, tt.args.enableIndentation, tt.args.outputFile, tt.args.rdr(), mockWrt)
			if !reflect.DeepEqual(mockWrt.WrittenSummary, tt.want) {
				t.Errorf("mergeAction() = %v, want %v", mockWrt.WrittenSummary, tt.want)
			}
		})
	}
}

type mockReader struct {
	ReadDirectoryFn func(path string, parser func(filePath string, data []byte)) error
}

func (r mockReader) ReadDirectory(path string, parser func(filePath string, data []byte)) error {
	return r.ReadDirectoryFn(path, parser)
}

type mockWriter struct {
	WrittenSummary board.Summary
	WriteFileFn    func(obj interface{}, filePath string, enableIndentation bool) error
}

func (w mockWriter) WriteFile(obj interface{}, filePath string, enableIndentation bool) error {
	return w.WriteFileFn(obj, filePath, enableIndentation)
}
