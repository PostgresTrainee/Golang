package main

import (
	"fmt"
)

const s string = "welcome"

func main() {
	fmt.Println(s)

	const n = 40000
	const d = 520m / n
	fmt.Println(d)
}