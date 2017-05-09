package rediscli

import "errors"

// CreateChannel find or create channel in memory
func CreateChannel(merchantID string, uuid string) (*Channel, error) {

	if len(merchantID) == 0 {
		return nil, errors.New("merchant_id not found")
	}

	if len(uuid) == 0 {
		return nil, errors.New("channel_id not found")
	}

	ch, err := client.createChannel(merchantID, uuid)

	return ch, err
}
