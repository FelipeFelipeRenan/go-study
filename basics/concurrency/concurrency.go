package main

import (
	"fmt"
	"time"
)

// Goroutine are a concurrent way to run code

func main() {

	// Creating a goroutine using a function
	go minhaFunc()

	// Waiting a time to execute the goroutine
	time.Sleep(time.Second)
	fmt.Println("Programa principal encerrado")

	ch := make(chan int)

	go func ()  {
		ch <- 42
	}()

	valor := <-ch
	fmt.Println("Valor recebido: ", valor)
}

func minhaFunc(){
	fmt.Println("Executanto uma goroutine")
}