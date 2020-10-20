package main

import (
	"log"
	"time"
)

func producer(q chan int, c chan int) {
	ticker := time.NewTicker(time.Second / 2)
	defer ticker.Stop()

	for i := 10; i < 10000; i++ {
		<-ticker.C
		select {
		case <-c:
			log.Println("Producer stopped")
			return
		default:
			q <- i
			log.Printf("Produce %d\n", i)
		}
	}
}

func consumer(q chan int, c chan int) {
	ticker := time.NewTicker(time.Second / 10)
	defer ticker.Stop()

	for {
		<-ticker.C
		select {
		case v := <-q:
			log.Printf("Consume %d\n", v)
		case <-c:
			log.Println("Consumer stopped")
		default:
			log.Println("................ Consumer is blocked")
		}
	}
}

func controller(c chan int) {
	time.Sleep(8 * time.Second)
	close(c)
}

func main() {
	q := make(chan int, 100)
	c := make(chan int)

	log.Println("#1 Start Controller")
	go controller(c)

	log.Println("#2 Start Consumer")
	go consumer(q, c)
	time.Sleep(1 * time.Second)

	log.Println("#3 Start Producer")
	go producer(q, c)

	time.Sleep(10 * time.Second)
}
