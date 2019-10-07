package main

import (
	"fmt"
	"time"
)

// Sender for get message from channel
type Sender chan<- int

// Receiver for send message to channel
type Receiver <-chan int

func main3() {

	fmt.Printf("starting...\n")

	myChannel := make(chan int, 1)
	go func() {
		var sender Sender = myChannel
		sender <- 1
		fmt.Printf("sent!\n")
	}()

	go func() {
		
		var receiver Receiver = myChannel
		time.Sleep(2*time.Second)
		receivedValue := <-receiver
		fmt.Printf("received! %v\n", receivedValue)
	}()

	time.Sleep(3*time.Second)
}
