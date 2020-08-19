package main

import (
	"fmt"
	"time"
)

func consumer(cname string, ch chan int) {
	for i := range ch {
		fmt.Println(fmt.Sprintf("Consumer -- %s: %d", cname, i))
	}
	fmt.Println("ch closed.")
}

func producer(pname string, ch chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println(fmt.Sprintf("procuder -- %s: %d", pname, i))
		ch <- i
	}
}

func main() {
	data := make(chan int)

	go producer("producer1", data)
	go producer("producer2", data)
	go consumer("consumer1", data)
	go consumer("consumer2", data)

	time.Sleep(18 * time.Second)
	close(data)
	time.Sleep(18 * time.Second)
}
