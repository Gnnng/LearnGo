package main

import (
	"fmt"
	"time"
)

var c chan int

func consumer(t int) {
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("I am a consumer")
	c <- 1
}

func producer(t int) {
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("I am a producer")
	c <- 1
}

func main() {
	c = make(chan int)
	go consumer(20)
	go producer(10)
	//time.Sleep(time.Duration(5) * time.Second)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
