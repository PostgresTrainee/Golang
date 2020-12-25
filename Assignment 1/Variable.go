package main

import "fmt"

var a = 106

func main() {
	b := 111
	fmt.Println(a)
	fmt.Println(b)

	foo()
}

func foo() {
	fmt.Println("Desktop", b)
}