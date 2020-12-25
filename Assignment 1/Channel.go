package main
import (
	"fmt"
)
func main()
{
	myChannel := make(chan string)
	go func()
	{
		myChannel <-"Hello"
		myChannel <- "World"
	}
	text := <- mychannel
	fmt.Println(text)
	fmt.Println(<- myChannel)
}
