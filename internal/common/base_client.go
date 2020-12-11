package common

import (
	"fmt"
	"log"

	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
)

const id = 420

func NewClient(id shared.ClientID) Client {
	return &BaseClient{id: id}
}

type BaseClient struct {
	id shared.ClientID
}

func (c *BaseClient) Echo(s string) string {
	c.logf("Echo: '%v'", s)
	return s
}

func (c *BaseClient) GetID() shared.ClientID {
	return c.id
}

// logf is the client's logger that prepends logs with your ID. This makes
// it easier to read logs. DO NOT use other loggers that will mess logs up!
func (c *BaseClient) logf(format string, a ...interface{}) {
	log.Printf("[%v]: %v", c.id, fmt.Sprintf(format, a...))
}

func (c *BaseClient) StartOfTurnUpdate(gameState GameState) {
	c.logf("Received game state update: %v", gameState)
	// TODO
}

func (c *BaseClient) EndOfTurnActions() []Action {
	c.logf("EndOfTurnActions")
	return nil
}
