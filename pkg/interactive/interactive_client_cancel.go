package interactive

import (
	"context"
	"log"
)

// create a cancel context for the interactive prompt, and set c.cancelFunc
func (c *InteractiveClient) createPromptContext(parentContext context.Context) context.Context {
	// ensure previous prompt is cleaned up
	if c.cancelPrompt != nil {
		c.cancelPrompt()
	}
	ctx, cancel := context.WithCancel(parentContext)
	c.cancelPrompt = cancel
	return ctx
}

func (c *InteractiveClient) createQueryContext(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	c.cancelActiveQuery = cancel
	return ctx
}

func (c *InteractiveClient) cancelActiveQueryIfAny() {
	if c.cancelActiveQuery != nil {
		log.Println("[INFO] cancelActiveQueryIfAny CALLING cancelActiveQuery")
		c.cancelActiveQuery()
		c.cancelActiveQuery = nil
	} else {
		log.Println("[INFO] cancelActiveQueryIfAny NO active query")
	}
}
