package cmd

import (
	"envchecker/config"
	"envchecker/g"
	"envchecker/utils"
	"os"

	"github.com/urfave/cli/v2"
)

func app() *cli.App {
	var confPath string
	var commands []*cli.Command
	commands = append(commands, cliDownloader())
	commands = append(commands, cliArchiver()...)
	app := &cli.App{
		Name: "envchecker",
		Authors: []*cli.Author{
			{
				Name: "ws",
			},
		},
		Version: "0.0.1",
		Usage:   "detect the necessary dependencies in the system environment",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "load configuration from `FILE`",
				Value:       "envchecker.json",
				Destination: &confPath,
			},
		},
		Commands: commands,
		Action: func(*cli.Context) error {
			if utils.IsPathExist(confPath) {
				c, err := config.From(confPath)
				if err != nil {
					g.Log.Error("config.From", err)
					return err
				}
				err = c.Check()
				if err != nil {
					g.Log.Error("config.Check", err)
					return err
				}
			} else {
				if err := config.Init(); err != nil {
					g.Log.Error("config.Init", err)
					return err
				}
				g.Log.Info("generate config finished, please restart")
			}
			return nil
		},
	}
	return app
}

func Executor() error {
	return app().Run(os.Args)
}
