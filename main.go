package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "keil",
		Usage: "Manages development boards on Arm Keil system",
		Commands: []*cli.Command{
			{
				Name:        "merge",
				Usage:       "Merges boards metadata from json files into one file",
				Description: "Merges boards metadata from json files into one file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "dir",
						Usage:    "Single directory with all files to be merged",
						Aliases:  []string{"d"},
						OnlyOnce: true,
						Required: true,
					},
				},
				Before: func(ctx context.Context, command *cli.Command) (context.Context, error) {
					return ctx, nil
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					return nil
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
