package main

import (
	"code/internal/helpers"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   true,
				Usage:   "show total size of directory including all subdirectories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Value:   false,
				Aliases: []string{"h"},
				Usage:   "print sizes in human-readable format",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "show sizes for all files including hidden files",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			isRecursive := cmd.Bool("recursive")
			isHumanReadable := cmd.Bool("human")
			shouldShowAll := cmd.Bool("all")

			size, err := helpers.GetPathSize(path, isRecursive, isHumanReadable, shouldShowAll)
			if err != nil {
				return err
			}
			fmt.Println(size)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
