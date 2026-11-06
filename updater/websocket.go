package updater

import (
	"context"
	"log"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type GetType struct {
	Type string `json:"type"`
}

// TODO
func openWebsocket() {
	ctx := context.Background()

	c, _, err := websocket.Dial(ctx, "wss://sae.tib.org/saews/public-events/11417623", nil)
	if err != nil {
		// ...
	}
	defer c.CloseNow()

	for {
		var jsonType GetType

		err = wsjson.Read(ctx, c, &jsonType)
		if err != nil {
			continue
		}

		switch jsonType.Type {
		case "position":
			println("Getting position")
		case "esta-info":
			println("Getting bus info")
		case "stop":
			println("Arriving to stop")
		case "close":
			println("Closing this websocket")
			err = c.Close(websocket.StatusNormalClosure, "")
			if err != nil {
				break
			}
		default:
			log.Println("Unknown type:", jsonType.Type)
		}
	}
}
