package main

import "fmt"

func calcularFatorial(n int) int {
	resultado := 1
	for i := 1; i < n; i++ {
		resultado *= i
	}
	return resultado
}
func main() {
	numeros := []int{5, 10, 15}

	// creaing a channel to get the values from factoral
	resultados := make(chan int)

	for _, value := range numeros {
		// creaing am amonymouss function using goroutine to compute the factorial
		// and pass the value to resoltados channel 
		go func(n int) {
			resultados <- calcularFatorial(n)
		}(value)

	}

	for range numeros {
		fmt.Println("Fatorial: ", <-resultados)
	}
}
