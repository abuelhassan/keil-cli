package board

import (
	"reflect"
	"testing"
)

func TestSummary_AppendBoards(t *testing.T) {
	type fields struct {
		Boards   []Board
		Metadata Metadata
	}
	type args struct {
		boards []Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Summary
	}{
		{
			name: "Should sort boards and create metadata",
			fields: fields{
				Boards:   make([]Board, 0),
				Metadata: Metadata{},
			},
			args: args{
				boards: []Board{
					{
						Name:   "B7-400X",
						Vendor: "Boards R Us",
					},
					{
						Name:   "Low_Power",
						Vendor: "Tech Corp.",
					},
					{
						Name:   "D4-200S",
						Vendor: "Boards R Us",
					},
				},
			},
			want: Summary{
				Boards: []Board{
					{
						Name:   "B7-400X",
						Vendor: "Boards R Us",
					},
					{
						Name:   "D4-200S",
						Vendor: "Boards R Us",
					},
					{
						Name:   "Low_Power",
						Vendor: "Tech Corp.",
					},
				},
				Metadata: Metadata{
					TotalVendors: 2,
					TotalBoards:  3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Summary{
				Boards:   tt.fields.Boards,
				Metadata: tt.fields.Metadata,
			}
			s.AppendBoards(tt.args.boards)
			if !reflect.DeepEqual(s.Metadata, tt.want.Metadata) {
				t.Errorf("Metadata = %v, want %v", s.Metadata, tt.want.Metadata)
			}
			if !reflect.DeepEqual(s.Boards, tt.want.Boards) {
				t.Errorf("Boards = %v, want %v", s.Boards, tt.want.Boards)
			}
		})
	}
}
