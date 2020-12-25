package main 
import (
	"fmt"
	"time"
)
func main()
{
	go myFunc("First")
	go myFunc("Second")
	go myFunc("Third")
	fmt.Println("done")
	time.Sleep(time.Second)
	fmt.Println("done 2")
}
func myFunc (a string)
{
	for i:=0; i<3; i++
	{
		fmt.Println(a, ":", i)
	}
}