package main

import (
	"log"
	"os"

	"path/filepath"

	converter "github.com/Angelmaneuver/belt-conveyor/internal/converter/webp"
	"github.com/Angelmaneuver/belt-conveyor/internal/manager"
	"github.com/Angelmaneuver/belt-conveyor/internal/watcher"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "Belt Conveyor ver.Webp",
		Usage:     "Converts image files placed on the watch point to webp.",
		UsageText: "go run cmd/webp.go [global options] command [command options]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "watchpoint",
				Aliases:  []string{"wp"},
				Usage:    "`Directory Path` to watch.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "destination",
				Aliases:  []string{"d"},
				Usage:    "`Directory Path` to store conversion result files.",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "quality",
				Aliases:  []string{"q"},
				Usage:    "For WEBP, it can be a quality from 1 to 100 (the higher is the better). By default (without any parameter) and for quality above 100 the lossless compression is used.",
				Required: false,
				Value:    100,
			},
		},
		Action: func(ctx *cli.Context) error {
			var watchpoint = ctx.String("wp")
			var destination = ctx.String("d")
			var quality = ctx.Int("q")

			if watchpoint == destination {
				log.Fatal("Error watchpoint and destination can't specify the same location")
			}

			option := converter.NewOption(quality)

			c, err := converter.New(option)
			if err != nil {
				log.Fatal("Error ", err)
			}

			var converter manager.Converter = c

			manager, err := manager.New(&converter, watchpoint, destination)
			if err != nil {
				log.Fatal("Error ", err)
			}

			filepath.Walk(watchpoint, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Println("Error ", err)
					return err
				}

				if info.IsDir() {
					return nil
				}

				if err := manager.Run(path); err != nil {
					log.Println("Error ", err)
				}

				return nil
			})

			watcher.Start(watchpoint, manager)

			return nil
		},
	}

	app.Run(os.Args)
}
