package rediscli

import "encoding/json"

// Pubsub exported
func Pubsub(channelUUID string, payload interface{}) {

	response, _ := json.Marshal(payload)

	client.rds.Publish(channelUUID, string(response))

}
