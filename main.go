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

type outputtable interface {
	Display()
	saver
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

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
	printSomething(1234);
	printSomething(12.34);
	printSomething(true);
	printSomething("Hello, World!");
}

func printSomething(value interface{}) {
	fmt.Println(value)
}

func outputData (data outputtable) error {
	data.Display()
	err := saveData(data);
	return err
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