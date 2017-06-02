package caching

// DeleteOrder exported
func (c *RedisCache) DeleteOrder(merchant string, uuid string) error {

	_, err := c.getClient().HDel(merchant, uuid).Result()

	return err
}
