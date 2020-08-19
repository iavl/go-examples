package main

import (
	"time"

	PC "github.com/iavl/go-examples/channel"
)

func main() {
	data := make(chan int)

	go PC.Producer("producer1", data)
	go PC.Producer("producer2", data)
	go PC.Consumer("consumer1", data)
	go PC.Consumer("consumer2", data)

	time.Sleep(18 * time.Second)
	close(data)
	time.Sleep(18 * time.Second)

}
