package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

// Note ...
type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// AddNote ...
func AddNote(title, body string) {
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
func ReadNote(title string) (Note, error) {
	notes := loadNotes()
	for _, note := range notes {
		if note.Title == title {
			return note, nil
		}
	}
	return Note{}, errors.New("No note found")
}

// RemoveNote ...
func RemoveNote(title string) {

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
