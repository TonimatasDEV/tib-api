package updater

import (
	"context"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

// TODO
func openWebsocket() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:8080", nil)
	if err != nil {
		// ...
	}
	defer c.CloseNow()

	err = wsjson.Read(ctx, c, "hi")
	if err != nil {

	}

	c.Close(websocket.StatusNormalClosure, "")
}
