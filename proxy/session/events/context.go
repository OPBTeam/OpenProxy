package events

import "go.uber.org/atomic"

type Context struct {
	cancelled *atomic.Bool
}

func NewContext() *Context {
	return &Context{
		cancelled: atomic.NewBool(false),
	}
}

func (c *Context) Cancel() {
	c.cancelled.Store(true)
}
func (c *Context) IsCancelled() bool {
	return c.cancelled.Load()
}
