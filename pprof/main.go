package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	go func() {
		for {
			fmt.Println("running...")
			time.Sleep(10 * time.Millisecond)
		}
	}()
	select {}
}
