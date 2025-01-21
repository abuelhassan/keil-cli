package writer

import (
	"testing"
)

func Test_writer_WriteFile(t *testing.T) {
	type args struct {
		obj               interface{}
		filePath          string
		enableIndentation bool
	}
	var tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: The happy path can be better tested with buffered writing,
		// TODO: as the interface for buffered writing can receive a test stream.
		{
			name: "Should return an error in case of invalid path",
			args: args{
				obj:               nil,
				filePath:          "",
				enableIndentation: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &writer{}
			if err := w.WriteFile(tt.args.obj, tt.args.filePath, tt.args.enableIndentation); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
