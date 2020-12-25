package main
import "fmt"
func main()
{
	x := max(5,10)
	fmt.Println(x)
}
func max(x int,y int)int
{
	if(x>y)
	{
		return x
	}
	else
	{
		return y
	}
}