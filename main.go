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
					rdr, wrt := reader.New(), writer.New()
					dir := command.String(flagDirectory)
					enableIndentation := command.Bool(flagEnableIndentation)
					outputFile := command.String(flagOutput)
					if outputFile == "" {
						outputFile = defaultOutputFile
					}

					mergeAction(dir, enableIndentation, outputFile, rdr, wrt)
					return nil
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func mergeAction(dir string, enableIndentation bool, outputFile string, rdr reader.Reader, wrt writer.Writer) {
	boards := make([]board.Board, 0)
	err := rdr.ReadDirectory(dir, func(filePath string, data []byte) {
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

	err = wrt.WriteFile(summary, outputFile, enableIndentation)
	if err != nil {
		log.Fatal(err)
	}
}
