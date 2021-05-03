package main

import "fmt"

type bot interface {
  getGreeting() string
}

type englishBot struct{}
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

func (englishBot) getGreeting() string {
  /*
    Lets pretend that the implementation for this function
    will massively vary from the implemtentation of getGreeting for spanisBot
  */
  return "Hi there!"
}

func (spanishBot) getGreeting() string {
  return "Hola!"
}

