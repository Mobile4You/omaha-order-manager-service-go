package rediscli

import (
	"errors"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// SubPub terminal
func SubPub(channelUUID string, merchantUUID string, logicNumber string) (models.Terminal, Channel, error) {

	var terminal models.Terminal
	var ch Channel
	var err error
	ch, err = validadeSub(channelUUID, merchantUUID)
	if err != nil {
		return terminal, ch, err
	}

	terminal, ok := ch.Terminals[logicNumber]
	if !ok {

		terminal = models.Terminal{
			Number: logicNumber,
		}

		ch.Terminals[logicNumber] = terminal
	}

	return terminal, ch, nil
}

func validadeSub(channelUUID string, merchantUUID string) (Channel, error) {

	ch, err := findChannel(channelUUID)
	if err != nil {
		return *ch, errors.New("channel_id not found")
	}

	if ch.MerchantID != merchantUUID {
		return *ch, errors.New("merchant_id not autorized")
	}

	return *ch, nil
}
