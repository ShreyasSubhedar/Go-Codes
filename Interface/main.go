package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}
type spansihBot struct {
}

func main() {
	eb := englishBot{}
	sb := spansihBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// some dangerous logic to get ...
	return "Hi There!!"
}

func (spansihBot) getGreeting() string {
	// some very dangerous logic to get:
	return "Hola!!"
}
