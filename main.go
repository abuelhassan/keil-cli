package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

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

func main() {
	cmd := &cli.Command{
		Name:  "keil",
		Usage: "Manages development boards on Arm Keil system",
		Commands: []*cli.Command{
			{
				Name:  "merge",
				Usage: "Merges boards metadata from json files into one file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "dir",
						Usage:    "Single directory with all files to be merged",
						Aliases:  []string{"d"},
						OnlyOnce: true,
						Required: true,
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					path := command.String("dir")
					c, err := readDirectory(path)
					if err != nil {
						log.Fatal(err)
					}

					err = writeFile("output.json", c)
					return nil
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func readDirectory(path string) (FileContent, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return FileContent{}, fmt.Errorf("couldn't read directory %s. %v", path, err)
	}

	merged := FileContent{}
	for _, e := range entries {
		data, err := os.ReadFile(filepath.Join(path, e.Name()))
		if err != nil {
			return FileContent{}, fmt.Errorf("couldn't read file %s. %v", e.Name(), err)
		}

		c := FileContent{}
		err = json.Unmarshal(data, &c)
		if err != nil {
			return FileContent{}, fmt.Errorf("couldn't parse file %s. %v", e.Name(), err)
		}

		merged.Boards = append(merged.Boards, c.Boards...)
	}

	return merged, nil
}

func writeFile(path string, c FileContent) error {
	// TODO: Indentation can be optional based on a flag
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't marshal output %v", err)
	}

	err = os.WriteFile("output.json", data, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("couldn't write file %s. %v", path, err)
	}
	return nil
}
