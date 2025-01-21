package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/abuelhassan/keil-cli/board"
	"github.com/abuelhassan/keil-cli/reader"
	"github.com/abuelhassan/keil-cli/writer"

	"github.com/urfave/cli/v3"
)

const (
	// Flag names
	flagDirectory         = "dir"
	flagOutput            = "out"
	flagEnableIndentation = "enableIndentation"

	defaultOutputFile = "out.json"
)

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
					&cli.BoolFlag{
						Name:     flagEnableIndentation,
						Usage:    "Enables Indentation in output",
						Value:    false,
						OnlyOnce: true,
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					path := command.String(flagDirectory)
					enableIndentation := command.Bool(flagEnableIndentation)
					boards := make([]board.Board, 0)
					err := reader.New().ReadDirectory(path, func(filePath string, data []byte) {
						c := board.Summary{}
						err := json.Unmarshal(data, &c)
						if err != nil {
							log.Fatalf("couldn't parse file %s. %v", filePath, err)
						}
						boards = append(boards, c.Boards...)
					})
					if err != nil {
						log.Fatal(err)
					}

					summary := board.Summary{}
					summary.AppendBoards(boards)

					outputFile := command.String(flagOutput)
					if outputFile == "" {
						outputFile = defaultOutputFile
					}
					err = writer.New().WriteFile(summary, outputFile, enableIndentation)
					if err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
