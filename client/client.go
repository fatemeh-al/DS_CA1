package client

import (
	"fmt"
	"log"
	"time"
	"github.com/fatemeh-al/DS_CA1/broker"
)

type Client struct{
	channelName string
	b broker.Broker
}

func NewClient(channelName string, b broker.Broker) *Client{
	return &Client{channelName: channelName, b:b}
}

func (c *Client) RecieveMessage() {

	// subCh is a readony channel that we will
	// receive messages published on "ch1".
	subCh, err := c.b.Subscribe(c.channelName)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(3 * time.Second)
	for m := range subCh {
		fmt.Printf("got message: %s\n", m)
		break
	}
}