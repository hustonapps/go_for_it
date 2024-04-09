package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hustonapps.com/go_for_it/note"
	"hustonapps.com/go_for_it/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData();
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error creating todo: ", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating note: ", err)
		return
	}

	todo.Display()
	err = saveData(todo)

	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed", err)
		return err
	}

	fmt.Println("Note saved successfully!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt);
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\n")

	return text;
}