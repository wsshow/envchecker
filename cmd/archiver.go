package cmd

import (
	"envchecker/pkg/archiver"
	"envchecker/pkg/pterm"
	"envchecker/utils"
	"fmt"

	"github.com/urfave/cli/v2"
)

func compress() *cli.Command {
	var srcPath, dstPath string
	return &cli.Command{
		Name:     "compress",
		Category: "archiver",
		Aliases:  []string{"zip"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "src",
				Aliases:     []string{"c"},
				Usage:       "the path to be compressed",
				Required:    true,
				Destination: &srcPath,
			},
			&cli.StringFlag{
				Name:        "dst",
				Aliases:     []string{"d"},
				Usage:       "the path of the compressed file",
				Value:       "./",
				Destination: &dstPath,
			},
		},
		Action: func(ctx *cli.Context) error {
			if !utils.IsPathExist(srcPath) {
				pterm.Warning(fmt.Sprintf("not found src path: %s", srcPath))
				return nil
			}
			if err := archiver.Compress([]string{srcPath}, dstPath); err != nil {
				pterm.Error(fmt.Sprintf("compress error: %s", err.Error()))
			}
			return nil
		},
	}
}

func decompress() *cli.Command {
	var srcPath, dstPath string
	return &cli.Command{
		Name:     "decompress",
		Category: "archiver",
		Aliases:  []string{"unzip"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "src",
				Aliases:     []string{"c"},
				Usage:       "the path of the file to be decompressed",
				Required:    true,
				Destination: &srcPath,
			},
			&cli.StringFlag{
				Name:        "dst",
				Aliases:     []string{"d"},
				Usage:       "the path of the decompressed folder",
				Value:       "./",
				Destination: &dstPath,
			},
		},
		Action: func(ctx *cli.Context) error {
			if err := archiver.Decompress(srcPath, dstPath); err != nil {
				pterm.Error(fmt.Sprintf("decompress error: %s", err.Error()))
			}
			return nil
		},
	}
}

func cliArchiver() []*cli.Command {
	return []*cli.Command{
		compress(), decompress(),
	}
}
