package main

import (
	"fmt"
)

type Facebook struct {
	firstName string
	lastName  string
	price     int
	quantity  int
}

func main() {
	fb := Facebook{
		firstName: "kowshick",
		lastName:  "ravi",
		Tags: "Frnds"    ,
		
	}
	fmt.Println("Facebook", fb)
}
