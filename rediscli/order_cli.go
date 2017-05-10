package rediscli

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/go-redis/redis"
)

type (

	// OrderClient is an exported
	OrderClient struct {
		rds      *redis.Client
		channels map[string]*Channel
		sync.RWMutex
	}

	// Channel is an exported
	Channel struct {
		UUID       string `json:"id"`
		MerchantID string
		Terminals  map[string]*models.Terminal
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		sync.Mutex
	}
)

func (c *OrderClient) putOrder(merchant string, number string, value string) (bool, error) {
	c.Lock()
	defer c.Unlock()
	result := c.rds.HSet(merchant, number, value)
	return result.Result()
}

func (c *OrderClient) delOrder(merchant string, number string) (int64, error) {
	c.Lock()
	defer c.Unlock()
	result := c.rds.HDel(merchant, number)
	return result.Result()
}

func (c *OrderClient) getOrder(merchantID string, number string) (*models.Order, error) {
	c.RLock()
	defer c.RUnlock()
	var o models.Order
	jsonOrder, _ := c.rds.HGet(merchantID, number).Result()
	err := json.Unmarshal([]byte(jsonOrder), &o)
	return &o, err
}

func (c *Channel) enterChannel(merchantID string, channelID string, logicNumber string) *models.Terminal {
	c.Lock()
	defer c.Unlock()
	t := &models.Terminal{
		Number: logicNumber,
		Sub:    client.rds.Subscribe(channelID),
	}
	c.Terminals[logicNumber] = t
	return t
}

func (c *OrderClient) getChannels(merchantID string) []*Channel {
	c.Lock()
	defer c.Unlock()
	var chs []*Channel
	for _, v := range c.channels {
		if v.MerchantID == merchantID {
			chs = append(chs, v)
		}
	}
	return chs
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
		ch = &Channel{
			UUID:       number,
			MerchantID: merchantID,
			Terminals:  make(map[string]*models.Terminal),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		c.channels[number] = ch
	}
	return ch, nil
}
