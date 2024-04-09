package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string
	Content string
	CreatedAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) ShowNote() {
	fmt.Printf("\nYour note titled\n\n%v\n\nhas the following content:\n\n%v\n\n", n.Title, n.Content)
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}