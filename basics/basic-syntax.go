package main

import "fmt"

func main() {
	// 3 types of variable declaration

	// var name type or var name type = value
	var nome string
	var idade int = 24

	// type infering
	idade2 := 25
	nome = "Felipe"

	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(idade2)

}