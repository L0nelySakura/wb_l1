package main

import (
	"fmt"
)


type Human struct {
	Name string
	Surname string
	Age int
}

func (h Human) greet() {
	fmt.Printf("Hello, my name is %s %s\n", h.Name, h.Surname)
}

// Наследовательский класс, также имеет метод greet()
type Action struct {
	Human 
	Hobby string
}

func (a Action) interests() {
	fmt.Printf("My hobby is %s\n", a.Hobby)
}

func main() {
	a := Action{
		Human: Human{
			Name: "Egor",
			Surname: "Petrov",
			Age: 21,
		},
		Hobby: "Programming",
	}
	a.greet()
	a.interests()
}
