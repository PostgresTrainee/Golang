package main
import "fmt"
func Fact(i int)int
{
	if(i<=1)
	{
		return 1;
	}
	return i * Fact(i-1)
}
func main()
{
	var i int = 5
	fmt.Printf("Factorial of %d is %d",i,Fact(i))
}