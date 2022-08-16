package main

import "fmt"

type bot interface {
	getGreeting() string
}

//type bot interface{getGreeting(string,int) (string,error)}
//(string,int) is list of argument types, (string,error) is list of return types
type englishBot struct {
}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

//concrete type vs interface type

//concrete type includes map, struct, int, string, englishBot
//can create a value out of directly and access and change
//interface type inlcude bot, cannot create directly
