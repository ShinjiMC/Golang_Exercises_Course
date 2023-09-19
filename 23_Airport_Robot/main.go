package main

import "fmt"

//Interface Greeter
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, guy Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", guy.LanguageName(), guy.Greet(name))
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

func main() {
	var p Italian
	fmt.Println("Test 1: ", SayHello("Luiggi", p))
	var i Portuguese
	fmt.Println("Test 2: ", SayHello("Ronaldihno", i))
}
