package main

import "fmt"

type englishBot struct{}
type spanishhBot struct{}


func main() {
  eb := englishBot{}
  sb := spanishhBot{}

  printGreeting(eb)
  printGreeting(sb)
}

func (englishBot) getGreeting() string {
  /*
    Lets pretend that the implementation for this function
    will massively vary from the implemtentation of getGreeting for spanisBot
  */
  return "Hi there!"
}

func (spanisBot) getGreeting() string {
  return "Hola!"
}

func printGreeting(eb englishBot) {
  fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
  fmt.Println(sb.getGreeting())
}
