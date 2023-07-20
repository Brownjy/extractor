package main

import (
	"context"
	"extractor/conf"
	"extractor/conf/grafana"
	"extractor/conf/serialize"
	"extractor/dep"
	"extractor/lib/util/filesys"
	"extractor/mongo"
	"fmt"
	"github.com/dtynn/dix"
	"github.com/filecoin-project/lotus/lib/lotuslog"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"os"
)

var logger = logging.Logger("main")

func main() {
	lotuslog.SetupLogLevels()

	commands := []*cli.Command{
		InitCmd(),
		GrafanaCmd(),
	}
	app := &cli.App{
		Name:                 "pando-dashboard",
		Usage:                "Extract data form mongo, present data through grafana",
		EnableBashCompletion: true,
		Flags:                []cli.Flag{},
		Commands:             commands,
	}

	app.Setup()

	if err := app.Run(os.Args); err != nil {
		logger.Errorf("cli error: %s", err)
		os.Exit(1)
	}
}

func InitCmd() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Initializes pando-dashboard config file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "repo-path",
				Usage:    "repository directory persistent pando-dashboard config file and data",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			repoPath := ctx.String("repo-path")
			cfgPath, err := conf.Filename(repoPath, "")
			if err != nil {
				fmt.Errorf("failed to get config path: %s", err)
			}
			cfgFileExists, err := filesys.IsFileExists(cfgPath)
			if err != nil {
				return fmt.Errorf("failure to check config file: %s", err)
			}
			if cfgFileExists {
				return fmt.Errorf("config file %s exists", cfgPath)
			} else {
				cfg := conf.Init()
				// write to file
				if err := serialize.WriteConfigFile(serialize.ConfigPath(cfgPath), cfg); err != nil {
					return nil
				}
			}
			return nil
		},
	}
}

func GrafanaCmd() *cli.Command {
	return &cli.Command{
		Name:  "grafana",
		Usage: "Extract data and Present data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "repo-path",
				Usage:    "grafana will read/persistent configs and data from the repository path",
				Required: false,
			},
		},
		Action: func(cCtx *cli.Context) error {
			repoPath := cCtx.String("repo-path")
			configPath, err := conf.Filename(repoPath, "")
			if err != nil {
				fmt.Errorf("failed to get config path: %s", err)
			}
			var ctx context.Context
			var components struct {
				fx.In
				Config  *conf.Config
				Storage *mongo.DB
				Grafana *grafana.Grafana
			}
			// di
			dix.New(ctx,
				dix.Override(new(serialize.ConfigPath), serialize.ConfigPath(configPath)),
				dep.Extractor(ctx, fxlog, &components),
			)

			// run一个grafana服务
			components.Grafana.Run(ctx)
			select {}
			return nil
		},
	}
}

var (
	log   = logging.Logger("bell")
	fxlog = &fxLogger{
		ZapEventLogger: log,
	}
)

type fxLogger struct {
	*logging.ZapEventLogger
}

// Printf impls fx.Printer.Printf
func (l *fxLogger) Printf(msg string, args ...interface{}) {
	l.ZapEventLogger.Debugf(msg, args...)
}
