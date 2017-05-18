package rediscli

import "github.com/arthurstockler/omaha-order-manager-service-go/models"

// UnSubPub terminal
func UnSubPub(terminal models.Terminal, ch Channel) error {

	delete(ch.Terminals, terminal.Number)

	err := terminal.Sub.Unsubscribe(ch.UUID)

	return err
}
