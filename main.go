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
	// time.Sleep(60 * time.Second)
	// for m := range channel {
	// 	fmt.Printf("got message: %s\n", m)
	// }
}

func main() {
	var b broker.Broker
	var c *client.Client

	// Trying the in-memory broker.
	b = broker.NewMemoryBroker()
	// b = broker.NewRedisBroker()

	// subCh is a readony channel that we will
	// receive messages published on "ch1".
	subCh, err := b.Subscribe("ch1")
	if err != nil {
		log.Fatalln(err)
	}

	c = client.NewClient("ch1")
	// start a publish loop
	// publish a message every second.
	go func() {
		defer b.Close()

		i := 0
		for {
			i++
			if err := b.Publish("ch1", fmt.Sprintf("message %d", i)); err != nil {
				log.Fatalln(err)
			}

			time.Sleep(time.Second)
			// stop after 5 iterations.
			if i == 5 {
				if err := b.Unsubscribe("ch1"); err != nil {
					log.Fatalln(err)
				}
				return
			}
		}
	}()
	
	c.recieveMessage()
	// read messages from subCh published on "ch1".
	// for m := range subCh {
	// 	fmt.Printf("got message: %s\n", m)
	// }
	recieveMessage(subCh)

}
