package usecase

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
	"github.com/go-redis/redis"
)

func subscribeChannel(w http.ResponseWriter, r *http.Request) {

	logicNumber := r.Header.Get("logic_number")
	merchantUUID := r.Header.Get("merchant_id")
	channelUUID := r.Header.Get("channel_id")

	log.Printf("Channel: %v", channelUUID)
	log.Printf("merchantUUID: %v", merchantUUID)
	log.Printf("number: %v", logicNumber)

	flusher, ok := w.(http.Flusher)
	if !ok {
		respondWithError(w, http.StatusInternalServerError, "Streaming unsupported!")
		return
	}

	terminal, ch, err := rediscli.SubPub(channelUUID, merchantUUID, logicNumber)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		err := rediscli.UnSubPub(terminal, ch)
		log.Printf("HTTP connection just closed: %v", err)
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	var receive error
	var msg *redis.Message
	for receive == nil {

		msg, receive = terminal.Sub.ReceiveMessage()

		fmt.Fprintf(w, "data: Message: %s\n\n", msg.Payload)

		flusher.Flush()
	}

}
