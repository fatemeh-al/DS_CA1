package client

import (
	"fmt"
	"log"
	"time"
	"github.com/fatemeh-al/DS_CA1/broker"
)

type Client struct{
	channelName string
}

func NewClient(channelName string) *Client{
	return &Client{channelName: channelName}
}

func (c *Client) recieveMessage() {
	var b broker.Broker

	// Trying the in-memory broker.
	b = broker.NewMemoryBroker()
	// b = broker.NewRedisBroker()

	// subCh is a readony channel that we will
	// receive messages published on "ch1".
	subCh, err := b.Subscribe(c.channelName)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(10 * time.Second)
	for m := range subCh {
		fmt.Printf("got message: %s\n", m)
	}
}