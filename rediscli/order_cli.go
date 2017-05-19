package rediscli

import (
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/go-redis/redis"
)

type (

	// OrderClient is an exported
	OrderClient struct {
		rds      *redis.Client
		channels map[string]Channel
		sync.RWMutex
	}

	// Channel is an exported
	Channel struct {
		UUID       string                     `json:"id"`
		MerchantID string                     `json:"merchant_id"`
		Terminals  map[string]models.Terminal `json:"terminals"`
		CreatedAt  time.Time                  `json:"created_at"`
		UpdatedAt  time.Time                  `json:"updated_at"`
		Conn       *redis.PubSub
	}
)

func (c *OrderClient) putOrder(merchant string, number string, value string) (bool, error) {
	result := c.rds.HSet(merchant, number, value)
	return result.Result()
}

func (c *OrderClient) delOrder(merchant string, number string) (int64, error) {
	result := c.rds.HDel(merchant, number)
	return result.Result()
}

func (c *OrderClient) getOrder(merchantID string, number string) (*models.Order, error) {
	var o models.Order
	jsonOrder, _ := c.rds.HGet(merchantID, number).Result()
	if len(strings.TrimSpace(jsonOrder)) == 0 {
		return nil, errors.New("order not found")
	}
	err := json.Unmarshal([]byte(jsonOrder), &o)
	return &o, err
}

func (c *OrderClient) getChannels(merchantID string) []Channel {
	c.Lock()
	defer c.Unlock()
	var ret = make([]Channel, 0)
	for _, ch := range c.channels {
		if ch.MerchantID == merchantID {
			ret = append(ret, ch)
		}
	}
	return ret
}

func (c *OrderClient) createChannel(merchantID string, number string) (*Channel, error) {
	_, err := c.getOrder(merchantID, number)
	if err != nil {
		return nil, errors.New("order not found")
	}

	c.Lock()
	defer c.Unlock()

	ch, ok := c.channels[number]
	if !ok {
		ch = Channel{
			UUID:       number,
			MerchantID: merchantID,
			Terminals:  make(map[string]models.Terminal),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Conn:       client.rds.Subscribe(number),
		}
		c.channels[number] = ch
	}
	return &ch, nil
}

func (c *OrderClient) showChannel(number string) Channel {
	c.Lock()
	defer c.Unlock()
	ch, _ := c.channels[number]
	return ch
}
