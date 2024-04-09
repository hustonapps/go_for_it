package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title string
	content string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}

func (n Note) ShowNote() {
	fmt.Printf("\nYour note titled\n\n%v\n\nhas the following content:\n\n%v\n\n", n.title, n.content)
}