package rediscli

import (
	"errors"
	"strings"
)

// FindChannel expoeted
func findChannel(channelUUID string) (*Channel, error) {

	if len(strings.TrimSpace(channelUUID)) == 0 {
		return nil, errors.New("channel_id not found")
	}

	ch := client.showChannel(channelUUID)

	if len(strings.TrimSpace(ch.UUID)) == 0 {
		return nil, errors.New("channel_id not found")
	}

	return &ch, nil
}
