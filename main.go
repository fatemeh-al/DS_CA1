package main

import (
	"fmt"
	"log"
	"time"
	"github.com/fatemeh-al/DS_CA1/broker"
)

type myMessage string

func recieveMessage(channel <-chan broker.Message){
	time.Sleep(6 * time.Second)
	for m := range channel {
		fmt.Printf("got message: %s\n", m)
		break
	}
}

func main() {
	var b broker.Broker

	// Trying the in-memory broker.
	b = broker.NewMemoryBroker()
	// b = broker.NewRedisBroker()

	// subCh is a readony channel that we will
	// receive messages published on "ch1".
	channel, err := b.CreateChannel("ch1")
	subscribedChannel, err2 := b.Subscribe("ch1")
	if err != nil {
		log.Fatalln(err)
	}
	if err2 != nil {
		log.Fatalln(err2)
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
	recieveMessage(channel)
	recieveMessage(subscribedChannel)
}
