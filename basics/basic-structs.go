package main

import "fmt"

// struct syntax just like C

type Pessoa struct {
	Nome  string
	Idade int
	Email string
}

// methods can me declared to structs and be called like:
// struct.method
func (p Pessoa) Apresentar() {
	fmt.Printf("Ola, meu nome é %s e tenho %d anos", p.Nome, p.Idade)
}

func (p Pessoa) isAdult() bool {
	if p.Idade > 18 {
		return true

	}
	return false
}

// interfaces are like struct but for methods
// they holds abstract methods to be implemented by structs
type Animal interface {
	SoundEmmit() string
}

type Cachorro struct {
	Nome string
}

// the type Cachorro implements the interface Animal, just like POO
func (c Cachorro) SoundEmmit() string {
	return "\nBarking!!"
}

func main() {
	pessoa := Pessoa{Nome: "Felipe", Idade: 24, Email: "teste@email.com"}

	fmt.Println(pessoa.Nome, pessoa.Idade)
	pessoa.Apresentar()

	var animal Animal
	animal = Cachorro{Nome: "Bolinha"}
	fmt.Println(animal.SoundEmmit())
}
