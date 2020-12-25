package main
import (
	"fmt"
)
func main()
{
	defer printText("First")
	defer printText("Second")
	printText("Hello")
	defer printText("third")
}	
func printText (p string)
{
	fmt.Println(p)
}
