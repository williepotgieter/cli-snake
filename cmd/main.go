package main

import (
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
					Name:  "size",
					Usage: "set board size for new game",
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
			Usage: "Displays current game info",
			Action: func(c *cli.Context) error {
				sg, err := pkg.GetGameState()
				if err != nil {
					return err
				}
				//fmt.Println("Saved game: ", *sg)
				sg.RenderBoard()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
