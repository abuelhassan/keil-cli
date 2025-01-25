package board

import (
	"slices"
)

type Board struct {
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Core    string `json:"core"`
	HasWifi bool   `json:"has_wifi"`
}

type Metadata struct {
	TotalVendors int `json:"total_vendors"`
	TotalBoards  int `json:"total_boards"`
}

// Summary contains the list of all appended boards with metadata
type Summary struct {
	Boards   []Board  `json:"boards"`
	Metadata Metadata `json:"_metadata"`
}

// AppendBoards appends board to the existing ones, and updates the metadata object.
// And sort all boards alphabetically by vendor then name.
func (s *Summary) AppendBoards(boards []Board) {
	s.Boards = append(s.Boards, boards...)
	vendors := make(map[string]struct{})
	for _, board := range s.Boards {
		vendors[board.Vendor] = struct{}{}
	}
	s.Metadata.TotalBoards = len(s.Boards)
	s.Metadata.TotalVendors = len(vendors)
	slices.SortFunc(s.Boards, func(a, b Board) int {
		if a.Vendor < b.Vendor {
			return -1
		}
		if a.Vendor > b.Vendor {
			return 1
		}
		if a.Name < b.Name {
			return -1
		}
		return 1
	})
}
