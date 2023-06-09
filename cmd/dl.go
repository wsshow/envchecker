package cmd

import (
	"envchecker/pkg/dl"
	"envchecker/pkg/pterm"
	"fmt"

	"github.com/urfave/cli/v2"
)

func cliDownloader() *cli.Command {
	return &cli.Command{
		Name:     "downloader",
		Aliases:  []string{"dl"},
		Category: "dl",
		Action: func(ctx *cli.Context) error {
			urls := ctx.Args().Slice()
			if len(urls) == 0 {
				pterm.Info("The current download list is empty")
				return nil
			}
			if pterm.Confirm("Whether to start downloading") {
				err := dl.BatchDownload(urls)
				if err != nil {
					pterm.Error(fmt.Sprintf("download failed: %s", err.Error()))
				}
			}
			return nil
		},
	}
}
