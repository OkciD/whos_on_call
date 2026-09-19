package sse

import (
	"github.com/google/uuid"
	"go.jetify.com/sse"
)

type Event sse.Event

func NewEventWithRawData(name string, rawData []byte, multiline bool) *Event {
	return &Event{
		ID:    uuid.NewString(),
		Event: name,
		Data:  sse.Raw(rawData),
		Split: multiline,
	}
}
