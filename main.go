package main

import (
	"log"
	"os"

	"github.com/chandanghosh/notes-app/notes"
)

func main() {

	noteUtils := &notes.NoteUtils{}

	cliApp := noteUtils.NewCliApp()

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
