package goidoit

import "context"

type ConsoleService service

type ListCommandWrapper struct {
	Success bool     `json:"success"`
	Lines   []string `json:"output"`
}

func (c *ConsoleService) ListCommands(ctx context.Context) (ListCommandWrapper, error) {
	return parse[ListCommandWrapper](c.client.Request(ctx, "console.commands.listCommands", nil))
}
