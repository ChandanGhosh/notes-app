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
	Notes []Note
}

// NewCliApp ...
func (n *NoteUtils) NewCliApp() *cli.App {
	cliApp := cli.NewApp()
	cliApp.Name = "cli based note app"
	cliApp.Version = "1.0.0"
	cliApp.EnableBashCompletion = true
	cliApp.Commands = n.getCommands()
	n.Notes = loadNotes()
	return cliApp
}

// getCommands ...
func (n *NoteUtils) getCommands() []*cli.Command {
	return []*cli.Command{

		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list notes",
			Action: func(c *cli.Context) error {
				return n.listNotes()

			},
		},
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
				return n.addNote(c.String("title"), c.String("body"))

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
				return n.readNote(c.String("title"))
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "remove a note",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "title",
					Usage:   "note title",
					Aliases: []string{"t"},
				},
			},
			Action: func(c *cli.Context) error {
				return n.removeNote(c.String("title"))
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update a note",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "title",
					Usage:   "note title",
					Aliases: []string{"t"},
				},
				&cli.StringFlag{
					Name:    "body",
					Usage:   "note body",
					Aliases: []string{"b"},
				},
			},
			Action: func(c *cli.Context) error {
				return n.updateNote(c.String("title"), c.String("body"))
			},
		},
	}
}

// listNotes ...
func (n *NoteUtils) listNotes() error {
	n.Notes = loadNotes()
	if len(n.Notes) == 0 {
		return errors.New("no notes found")
	}
	for _, note := range n.Notes {
		fmt.Printf("%s:\t%s\n", note.Title, note.Body)
	}
	return nil
}

// addNote ...
func (n *NoteUtils) addNote(title, body string) error {
	notes := n.Notes
	for _, note := range notes {
		if note.Title == title {
			return errors.New("note title already taken")
		}
	}
	notes = append(notes, Note{Title: title, Body: body})
	n.saveNotes(notes)
	fmt.Println("New note added!")
	return nil
}

// ReadNote ...
func (n *NoteUtils) readNote(title string) error {
	notes := n.Notes
	for _, note := range notes {
		if note.Title == title {
			fmt.Println(note.Body)
			return nil
		}
	}
	return errors.New("No note found")
}

// RemoveNote ...
func (n *NoteUtils) removeNote(title string) error {
	notes := n.Notes
	noteIndex := getNoteIndex(title)
	if noteIndex == -1 {
		return errors.New("No notes found")
	}
	notes = append(notes[:noteIndex], notes[noteIndex+1:]...)
	n.saveNotes(notes)
	fmt.Println("Note removed!")
	return nil
}

func getNoteIndex(title string) int {
	for index, note := range loadNotes() {
		if note.Title == title {
			return index
		}
	}
	return -1
}

func (n *NoteUtils) updateNote(title, body string) error {
	noteIndex := getNoteIndex(title)
	if noteIndex == -1 {
		return errors.New("No note found to update")
	}

	n.Notes[noteIndex].Body = body
	n.saveNotes(n.Notes)
	fmt.Println("Note updated!")
	return nil
}

func (n *NoteUtils) saveNotes(notes []Note) {
	data, err := json.Marshal(notes)
	if err != nil {
		log.Fatalln("Error marshaling data", notes)
		return
	}
	err = ioutil.WriteFile("notes.json", data, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	n.Notes = loadNotes()

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
