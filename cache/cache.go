package cache

import (
	"encoding/json"
	"fmt"
)

type Cache struct {
	c map[string]interface{}
	write chan []byte
	read chan map[string]interface{}
	stop chan struct{}
}

func New() *Cache {
	return &Cache{
		c: make(map[string]interface{}),
		write: make(chan []byte),
		read: make(chan struct{}),
		stop: make(chan struct{}),
	}
}

// pass Cache so it knows it's using it
func(c *Cache) Start() {
	go c.start()
}

func(c *Cache) Stop() {
	// write to this channel
	c.stop <- struct{}{}
}

// lowercase private version
func(c *Cache) start() {
	for {
		// event loop router
		select {
		// when c.write triggers, make new var 'data'
		// out of data
		// c.write is channel here, if you're on left side of channel ur reading from it,
		// on other side, lik read, you're writing to it
		case data := <-c.write:
			c.handleWrite(data)

		// if reader is able to read, pass empty struct,
		// basically like passing null/nil
		case c.read <- struct{}{}:
			continue
		// when i get anthing, don't care what over the
		// stop channel, do this action

		// reading nothing in, just want to know that stop function
		// was called in any other channel 
		case <-c.stop:
			return
		}

	}
}

func(c *Cache) Write(d []byte) {
	c.write <- d
}

func(c *Cache) handleWrite(d []byte) {
	err := json.Unmarshal(d, &c.c)

	if err != nil {
		fmt.Printf("failed to parse data: %s", err)
	}
}


func(c *Cache) Read() map[string]interface{} {
	data := <-c.read
	return data
}