package board

import "sort"

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

type Summary struct {
	// Boards sort boards by vendor then by name
	Boards   []Board  `json:"boards"`
	Metadata Metadata `json:"_metadata"`
}

// AppendBoards appends board to the existing ones, and updates the metadata object.
func (s *Summary) AppendBoards(boards []Board) {
	s.Boards = append(s.Boards, boards...)
	vendors := make(map[string]struct{})
	for _, board := range s.Boards {
		vendors[board.Vendor] = struct{}{}
	}
	s.Metadata.TotalBoards = len(s.Boards)
	s.Metadata.TotalVendors = len(vendors)
	sort.Sort(s)
}

func (s *Summary) Len() int      { return len(s.Boards) }
func (s *Summary) Swap(i, j int) { s.Boards[i], s.Boards[j] = s.Boards[j], s.Boards[i] }
func (s *Summary) Less(i, j int) bool {
	if s.Boards[i].Vendor == s.Boards[j].Vendor {
		return s.Boards[i].Name < s.Boards[j].Name
	}
	return s.Boards[i].Vendor < s.Boards[j].Vendor
}
