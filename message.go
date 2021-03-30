package polis

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID
	SerialID uint64
	Metadata map[string]interface{}
	Body     []byte

	AuthorID  uuid.UUID
	ThreadID  *uuid.UUID
	Edited    bool
	CreatedAt time.Time
}

func (m Message) PlainText() string {
	return string(m.Body)
}

func NewMessage(body string, metadata map[string]interface{}) (Message, error) {
	return Message{}, nil
}

func (m Message) Update(body *string, metadata *map[string]interface{}) (Message, error) {
	return Message{}, nil
}

func (m Message) Delete() error {
	return nil
}
