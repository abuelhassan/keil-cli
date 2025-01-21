package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sort"

	"github.com/abuelhassan/keil-cli/reader"

	"github.com/urfave/cli/v3"
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

type FileContent struct {
	Boards   []Board  `json:"boards"`
	Metadata Metadata `json:"_metadata"`
}

type ByVendorThenName []Board

func (b ByVendorThenName) Len() int      { return len(b) }
func (b ByVendorThenName) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByVendorThenName) Less(i, j int) bool {
	if b[i].Vendor == b[j].Vendor {
		return b[i].Name < b[j].Name
	}
	return b[i].Vendor < b[j].Vendor
}

const (
	// Flag names
	flagDirectory = "dir"
	flagOutput    = "out"

	defaultOutputFile = "out.json"
)

func main() {
	rdr := reader.New()
	cmd := &cli.Command{
		Name:  "keil",
		Usage: "Manages development boards on Arm Keil system",
		Commands: []*cli.Command{
			{
				Name:  "merge",
				Usage: "Merges boards metadata from json files into one file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     flagDirectory,
						Usage:    "Single directory with all files to be merged",
						Aliases:  []string{"d"},
						OnlyOnce: true,
						Required: true,
					},
					&cli.StringFlag{
						Name:     flagOutput,
						Usage:    "Output file name",
						Value:    defaultOutputFile,
						Aliases:  []string{"o"},
						OnlyOnce: true,
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					path := command.String(flagDirectory)
					boards := make([]Board, 0)
					err := rdr.ReadDirectory(path, func(filePath string, data []byte) {
						c := FileContent{}
						err := json.Unmarshal(data, &c)
						if err != nil {
							log.Fatalf("couldn't parse file %s. %v", filePath, err)
						}
						boards = append(boards, c.Boards...)
					})
					merged := FileContent{Boards: boards}
					if err != nil {
						log.Fatal(err)
					}

					sort.Sort(ByVendorThenName(merged.Boards))
					populateMetadata(&merged)

					outputFile := command.String(flagOutput)
					if outputFile == "" {
						outputFile = defaultOutputFile
					}
					err = writeFile(outputFile, merged)
					return nil
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func populateMetadata(c *FileContent) {
	vendorsMp := make(map[string]struct{})
	for _, board := range c.Boards {
		vendorsMp[board.Vendor] = struct{}{}
	}
	c.Metadata.TotalVendors = len(vendorsMp)
	c.Metadata.TotalBoards = len(c.Boards)
}

func writeFile(filename string, c FileContent) error {
	// TODO: Indentation can be optional based on a flag
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't marshal output %v", err)
	}

	err = os.WriteFile(filename, data, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("couldn't write file %s. %v", filename, err)
	}
	return nil
}
