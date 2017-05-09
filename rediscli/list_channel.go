package rediscli

// ListChannel returns all channels of an EC
func ListChannel(merchantID string) []*Channel {
	return client.getChannels(merchantID)
}
