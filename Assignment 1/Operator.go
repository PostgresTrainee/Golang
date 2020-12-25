package main

import "fmt"

func main() {
	x, y := 21, 22
	fmt.Println("x+y=", x+y)
	fmt.Println("x-y=", x-y)
	fmt.Println("x*y=", x*y)
	fmt.Println("x/y=", x/y)
	fmt.Println("x mod y=", x%y)

	isbool := true
	fake := false

	fmt.Println(isbool && fake)
	fmt.Println(isbool || fake)
	fmt.Println(!isbool)
}