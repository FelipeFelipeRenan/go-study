package main

import "fmt"

// creating factorial function to test goroutines
func calcularFatorial(n int) int {
	resultado := 1
	for i := 1; i < n; i++ {
		resultado *= i
	}
	return resultado
}

func main() {

	numeros := []int{5,10,15}
	// creating channel to get results from factorial  
	resultados := make(chan int)

	for _, value := range numeros {
		go func(n int)  {
			resultados <- calcularFatorial(n)
		}(value)
		
	}

	for range numeros{
		fmt.Println("Fatorial: ", <- resultados)
	}
}