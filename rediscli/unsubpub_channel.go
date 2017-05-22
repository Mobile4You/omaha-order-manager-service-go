package rediscli

import "github.com/arthurstockler/omaha-order-manager-service-go/models"

// UnSubPub terminal
func UnSubPub(terminal models.Terminal, ch Channel) {

	delete(ch.Terminals, terminal.Number)

	if len(ch.Terminals) == 0 {
		ch.Conn.Unsubscribe(ch.UUID)
		client.deleteChannel(ch.UUID)
	}

}
