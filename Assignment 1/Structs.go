package main
import "fmt"
type Car struct 
{
	model string
	make string
	cc int
} 
func main()
{
	var Car1 car
	Car1.make = "TATA"
	Car1.model = "Kia"
	Car1.cc = 1200 
	printCar(Car1)
}