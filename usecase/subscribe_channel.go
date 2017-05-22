package usecase

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
	"github.com/go-redis/redis"
)

func subscribeChannel(w http.ResponseWriter, r *http.Request) {

	logic := r.Header.Get("logic_number")
	merchant := r.Header.Get("merchant_id")
	channel := r.Header.Get("channel_id")

	log.Printf("Channel: %v", channel)
	log.Printf("merchantUUID: %v", merchant)
	log.Printf("number: %v", logic)

	flusher, ok := w.(http.Flusher)
	if !ok {
		respondWithError(w, http.StatusInternalServerError, "Streaming unsupported!")
		return
	}

	terminal, ch, err := rediscli.SubPub(channel, merchant, logic)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		rediscli.UnSubPub(terminal, ch)
		log.Printf("HTTP connection just closed: %v", err)
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	//get last message channel
	or, _ := rediscli.FindOrder(merchant, channel)
	event := 0

	var receive error
	var msg *redis.Message
	go rediscli.Pubsub(channel, &or)
	for receive == nil {

		msg, receive = ch.Conn.ReceiveMessage()
		if event == 0 {
			fmt.Fprintf(w, "data: %s%s\n\n", "o", msg.Payload)
			event++
		} else {
			fmt.Fprintf(w, "data: %s%s\n\n", "i", msg.Payload)
		}

		flusher.Flush()
	}

}
