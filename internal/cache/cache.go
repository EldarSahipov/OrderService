package cache

import (
	"OrderService/internal/models"
	"sync"
)

type Cache struct {
	sync.RWMutex
	items map[string]Item
}

type Item struct {
	Value models.Order
}

func NewCache() *Cache {
	items := make(map[string]Item)

	cache := Cache{
		items: items,
	}

	return &cache
}

func (c *Cache) SetOrderFromCache(key string, order models.Order) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = Item{
		Value: order,
	}
}

func (c *Cache) GetOrderFromCacheById(key string) (models.Order, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]

	if !found {
		return models.Order{}, false
	}

	return item.Value, true
}
