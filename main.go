package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chandanghosh/notes-app/notes"
	"github.com/urfave/cli/v2"
)

func main() {

	cliApp := cli.NewApp()
	cliApp.Name = "cli based note app"
	cliApp.Version = "1.0.0"
	cliApp.EnableBashCompletion = true
	cliApp.Commands = []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a new note",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "title",
					Usage:   "Add a note title",
					Aliases: []string{"t"},
				},
				&cli.StringFlag{
					Name:    "body",
					Usage:   "Add a note body",
					Aliases: []string{"b"},
				},
			},
			Action: func(c *cli.Context) error {
				notes.AddNote(c.String("title"), c.String("body"))
				return nil
			},
		},
		{
			Name:    "read",
			Aliases: []string{"r"},
			Usage:   "read a note",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "title",
					Usage:   "Add a note title",
					Aliases: []string{"t"},
				},
			},
			Action: func(c *cli.Context) error {
				note, err := notes.ReadNote(c.String("title"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(note)
				return nil
			},
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
