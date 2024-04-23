package main

import (
	"fmt"
	"testing"
	"time"
)

// Go routine

func RunHelloWorld() {
	fmt.Println("Hello worlds")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()

	fmt.Println("end")

	time.Sleep(1 * time.Second)
}

func Counter(num int) {
	fmt.Println("Number is ", num)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go Counter(i)
	}

	time.Sleep(3 * time.Second)
}

// Channel

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "Hello channel"
		fmt.Println("Channel data sent")
	}()

	data := <-channel

	fmt.Println(data)
}

// Output
// Channel data sent
// Hello channel

// Channel In & Out

func OnlyIn(channel chan<- string) {
	channel <- "Only in"
}

func OnlyOut(channel <-chan string) {
	data := <-channel

	fmt.Println(data)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	defer close(channel)

	channel <- "Hello 1"
	channel <- "Hello 2"
	channel <- "Hello 3"
	channel <- "Hello 4"

	fmt.Println("finish")
}
