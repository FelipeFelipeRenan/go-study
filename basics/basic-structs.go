package main

import "fmt"

// struct syntax just like C

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	pessoa := Pessoa{Nome: "Felipe", Idade: 24}

	fmt.Println(pessoa.Nome, pessoa.Idade)
}