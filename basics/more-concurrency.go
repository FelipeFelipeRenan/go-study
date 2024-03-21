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
	numeros := []int{5,10,15}
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