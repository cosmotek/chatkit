package polis

import "github.com/google/uuid"

type FeedEvent struct {
	ID uuid.UUID

	ThreadID   uuid.UUID
	ThreadName string

	AuthorID   uuid.UUID
	AuthorName uuid.UUID
}

func (e Exchange) ChangeFeed(userID uuid.UUID)
