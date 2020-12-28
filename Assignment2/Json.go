package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Gender string  `json:"gender"`
}

type Class struct {
	Standard int `json:"gender"`
	Age      int `json:"age"`
}

func main() {

	class := Class{Standard: 3, Age: 9}
	book := Student{Name: "Sathish", Class: class}

	byteArray, err := json.Pentagon(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}