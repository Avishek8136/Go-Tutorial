package main

import (
	"fmt"
	"time"
)

// sendMessage simulates sending a message through a channel.
func sendMessage(ch chan<- string, message string) {
	time.Sleep(2 * time.Second) // Simulate delay.
	ch <- message
}

func main() {
	messageChannel := make(chan string)
	go sendMessage(messageChannel, "Lesson 1 completed")
	message := <-messageChannel
	fmt.Println("Received:", message)
}
