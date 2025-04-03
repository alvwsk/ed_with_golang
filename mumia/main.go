package main

import "fmt"

type Person struct {
	name     string
	age      int
	category string
}

func NewPerson(name string, age int) Person {
	person := Person{name: name, age: age}
	person.setCategory()
	return person
}

func (p *Person) setCategory() {
	switch {
	case p.age < 12:
		p.category = "crianca"
	case p.age < 18:
		p.category = "jovem"
	case p.age < 65:
		p.category = "adulto"
	case p.age < 1000:
		p.category = "idoso"
	default:
		p.category = "mumia"
	}
}

func (p Person) Exibir() string {
	return p.name + " eh " + p.category
}

func main() {
	var name string
	var age int

	fmt.Scanln(&name)
	fmt.Scanln(&age)

	person := NewPerson(name, age)
	fmt.Println(person.Exibir())
}
