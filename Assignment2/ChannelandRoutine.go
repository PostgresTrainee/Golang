package main
import (
	  "fmt"
	  "time"
)

func try(done chan bool)
{
	fmt.Println("Channel Started")
	time.Sleep(6 * time.Second)
	fmt.Println("Channel sleep for a while")
	done <- true
}
func main()
{
	done := make(chan bool)
	fmt.Println("Main Routine")
	go try(done)
	<-done
	fmt.Println("Received Data")
	
}