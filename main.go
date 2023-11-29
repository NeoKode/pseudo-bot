package main

import "fmt"

func helloWorld() string {
	return "Hola, mundo"
}

func main() {
	message := helloWorld()
	fmt.Println(message)
}
