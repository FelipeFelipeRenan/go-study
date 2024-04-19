package main

import "fmt"


// struct syntax just like C
// You can use embbeded structs to represent more complex types
type Endereco struct {
	Rua string
	CEP string
	Cidade string
}
type Pessoa struct {
	Nome  string
	Idade int
	Email string
	Endereco
}

// methods can me declared to structs and be called like:
// struct.method
func (p Pessoa) Apresentar() {
	fmt.Printf("Ola, meu nome é %s e tenho %d anos", p.Nome, p.Idade)
}

func (p Pessoa) IsAdult() bool {
	if p.Idade > 18 {
		return true

	}
	return false
}

// you can create methods assigned to pointers instead of types
// it allows the struct value mutation
func (p *Pessoa) GetOlder() int {
	return p.Idade + 1
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

// you can create interface with methods and "implements into other structs"
type FormaGeometricaa interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura float64
}

func (r Retangulo) Area() float64  {
	return r.Largura * r.Altura
}
func main() {
	pessoa := Pessoa{Nome: "Felipe", Idade: 24, Email: "teste@email.com"}
	fmt.Println("Created person")
	fmt.Println(pessoa.Nome, pessoa.Idade)
	fmt.Println("Updating age")
	fmt.Println(pessoa.GetOlder())
	fmt.Println("Presenting")
	pessoa.Apresentar()

	fmt.Println("\nInterface usage")
	var animal Animal
	animal = Cachorro{Nome: "Bolinha"}
	fmt.Println(animal.SoundEmmit())

	pessoa2 := Pessoa{
		Nome: "Maria",
		Idade: 25,
		Email: "maria@email.com",
		Endereco: Endereco{
			Rua: "Rua da rua",
			CEP: "123456-123",
			Cidade: "Aquela cidade",
		},
	}
	fmt.Println("Embbeded structs")
	fmt.Printf("%s mora em %s", pessoa2.Nome, pessoa2.Cidade)


	// interface methods
	retangulo := Retangulo{Largura: 10, Altura: 5}
	fmt.Println(retangulo.Area()) // Saída: 50.0



}
