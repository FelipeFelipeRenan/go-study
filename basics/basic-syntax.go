package main

import "fmt"


// functions works just like C
func soma(a int, b int) int {
	return a + b
}
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

	// conditionals

	if idade >= 24 {
		fmt.Println("Já passou o seu aniversario!")
	} else {
		fmt.Println("Ainda nao passou o seu aniversario!")
	}

	fmt.Println(soma(idade, idade2))

	// arrays 

	// the way how to declare static arrays
	var numbers [5]int

	numbers[0] = 10
	
	// slices are more flexibles than arrays
	numbers2 := []int{10, 20, 30, 40}

	fmt.Println(numbers)
	
	fmt.Println(numbers2)

	for i, value := range numbers2 {
		fmt.Println(i + 1, value)
		
	}

	// maps are key-value data types
	// they are declared as  myVar := map[type]type 
	
	myMap := make(map[string]int)

	myMap["Um"] = 1
	myMap["Dois"] = 2

	valor := myMap["Um"]

	fmt.Println(valor)
}
