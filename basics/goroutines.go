package main

import (
	"fmt"
	"time"
)

// function that creates the numbers to be consumed to another goroutine 
func produces(ch chan<- int) {
	// close channel afther the ending of the execution
	defer close(ch)

	// passing the values for iteration to channel ch
	for i := 1; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

// function that consumes the values from producer through a channel 
func consumes(ch <-chan int) {
	// consuming the channel values
	for num := range ch {
		fmt.Println("Consumindo número: ", num)
		time.Sleep(500 * time.Millisecond)

	}
}

func main() {
	// creating channel
	ch := make(chan int)

	// Adding the values of produces to channel ch
	go produces(ch)

	// taking the values of channel 
	go consumes(ch)

	time.Sleep(6*time.Second)

	fmt.Println("Programa principal encerrado!")
}
