package main

import (
	"fmt"
	"log"
	"time"
	"github.com/fatemeh-al/DS_CA1/broker"
	"github.com/fatemeh-al/DS_CA1/client"
)

type myMessage string

func recieveMessage(channel <-chan broker.Message){
	time.Sleep(3 * time.Second)
	for m := range channel {
		fmt.Printf("got message: %s\n", m)
		break
	}
}

func main() {
	var b broker.Broker
	var c1 *client.Client
	var c2 *client.Client

	// Trying the in-memory broker.
	b = broker.NewMemoryBroker()
	
	// subCh is a readony channel that we will
	// receive messages published on "ch1".
	channel, err := b.CreateChannel("ch1")
	c1 = client.NewClient("ch1", b)
	c2 = client.NewClient("ch1", b)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		defer b.Close()

		i := 0
		for {
			i++
			if err := b.Publish("ch1", fmt.Sprintf("message %d", i)); err != nil {
				log.Fatalln(err)
			}

			time.Sleep(time.Second)
			if i == 5 {
				if err := b.DeleteChannel("ch1"); err != nil {
					log.Fatalln(err)
				}
				return
			}
		}
	}()
	c1.RecieveMessage()
	c2.RecieveMessage()
	recieveMessage(channel)
}
