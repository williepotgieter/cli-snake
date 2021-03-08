package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/williepotgieter/cli-snake/pkg"
)

func main() {
	app := &cli.App{
		Name:  "CLI Snake",
		Usage: "Play snake in the command line",
	}

	app.Commands = []*cli.Command{
		{
			Name:  "newgame",
			Usage: "Creates a new game.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "size",
					Usage:       "set board size for new game",
					Value:       "10x22",
					DefaultText: "10x22",
				},
			},
			Action: func(c *cli.Context) error {
				bi, err := pkg.CreateBoard(c)
				if err != nil {
					return err
				}
				if err := bi.NewGame(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "info",
			Usage: "displays current game info",
			Action: func(c *cli.Context) error {
				sg, err := pkg.GetGameState()
				if err != nil {
					return err
				}
				sg.RenderBoard()
				return nil
			},
		},
		{
			Name:  "play",
			Usage: "Creates a new game.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "move",
					Usage:    "moves snake head one position (up, down, left or right)",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				sg, err := pkg.GetGameState()
				if err != nil {
					return errors.New("could not load game state. make sure to create a new game before running this command")
				}

				if err := sg.MoveSnake(c); err != nil {
					return err
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
