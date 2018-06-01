package main

import "fmt"

func main() {
	fmt.Printf(SayHello("john"))
}

func SayHello(name string) string {
	greeting := "hello, " + name
	return greeting
}
