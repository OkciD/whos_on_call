package sse

import (
	"context"
	"fmt"

	"go.jetify.com/sse"
)

type Conn struct {
	*sse.Conn

	ID string

	ctx      context.Context
	cancelFn context.CancelFunc
}

func (c *Conn) Context() context.Context {
	if c.ctx != nil {
		return c.ctx
	}

	return context.Background()
}

func (c *Conn) SendEvent(ctx context.Context, event *Event) error {
	e := sse.Event(*event)
	if err := c.Conn.SendEvent(ctx, &e); err != nil {
		return fmt.Errorf("error sending event: %w", err)
	}
	return nil
}
