package main

import (
	"fmt"
	"sync"
)

type CacheD struct {
	data  map[string]string
	mutex sync.RWMutex
}

var gCacheD CacheD

func init() {
	gCacheD.data = make(map[string]string)
}

func (c *CacheD) Add(key, value string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.data[key]; !exists {
		c.data[key] = value
		return nil
	}

	return fmt.Errorf("key %q already exists", key)
}

func (c *CacheD) Del(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
}

func (c *CacheD) Get(key string) (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if value, ok := c.data[key]; ok {
		return value, nil
	}

	return "", fmt.Errorf("key %+q does not exist", key)
}

func (c *CacheD) Update(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
}

func (c *CacheD) GetAll() (items [][2]string) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	items = make([][2]string, 0)

	for k, v := range c.data {
		items = append(items, [2]string{k, v})
	}

	return
}
