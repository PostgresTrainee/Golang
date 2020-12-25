package main
import "fmt"
import "math"
type Shape interface
{
	area() float64
}
type Circle struct
{
	radius() float64
}
func (c Circle) area() float64
{
	return math.Pi c.radius * c.radius
}
type Rectanlge struct
{
	height,width float64
}
func (r Rectanlge) area() float64
{
    return r.width * r.height
}
func getArea (s Shape) float64
{
	return s.area()
}
func main()
{
	c := Circle{radius : 10}
	r := Rectangle{width : 12, height : 14}
	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}