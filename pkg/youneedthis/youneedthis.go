package youneedthis

import (
	"fmt"
	"time"
)

// Message - public API for this awesomeness
type Message interface {
	fmt.Stringer
}

// internal implementation type
type message struct {
	Author      string
	Content     string
	PublishedAt time.Time
}

// NewMessage - you get the idea
func NewMessage(author, content string) Message {
	return &message{
		Author:      author,
		Content:     content,
		PublishedAt: time.Now().UTC(),
	}
}

func (m *message) String() string {
	return fmt.Sprintf("%s: %s (at: %s)", m.Author, m.Content, m.PublishedAt)
}
