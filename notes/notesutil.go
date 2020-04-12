package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/urfave/cli/v2"
)

// NoteUtils ...
type NoteUtils struct {
}

// NewCliApp ...
func (n *NoteUtils) NewCliApp() *cli.App {
	cliApp := cli.NewApp()
	cliApp.Name = "cli based note app"
	cliApp.Version = "1.0.0"
	cliApp.EnableBashCompletion = true

	cliApp.Commands = n.getCommands()
	return cliApp
}

// getCommands ...
func (n *NoteUtils) getCommands() []*cli.Command {
	return []*cli.Command{
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
				n.addNote(c.String("title"), c.String("body"))
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
				note, err := n.readNote(c.String("title"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(note)
				return nil
			},
		},
	}
}

// addNote ...
func (n *NoteUtils) addNote(title, body string) {
	notes := loadNotes()
	for _, note := range notes {
		if note.Title == title {
			fmt.Println("note tiitle already taken")
			return
		}
	}
	notes = append(notes, Note{Title: title, Body: body})
	saveNotes(notes)
	fmt.Println("New note added!")
}

// ReadNote ...
func (n *NoteUtils) readNote(title string) (Note, error) {
	notes := loadNotes()
	for _, note := range notes {
		if note.Title == title {
			return note, nil
		}
	}
	return Note{}, errors.New("No note found")
}

// RemoveNote ...
func (n *NoteUtils) removeNote(title string) {

}

func saveNotes(n []Note) {
	data, err := json.Marshal(n)
	if err != nil {
		log.Fatalln("Error marshaling data", n)

	}
	err = ioutil.WriteFile("notes.json", data, 0644)
	if err != nil {
		log.Println(err)
	}

}

func loadNotes() []Note {
	n := make([]Note, 0)
	data, err := ioutil.ReadFile("notes.json")
	if err != nil || len(data) == 0 {
		return n
	}

	err = json.Unmarshal(data, &n)
	if err != nil {
		log.Fatalln("error decoding json")
	}
	return n
}
