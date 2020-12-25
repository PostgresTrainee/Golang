package main
import "fmt"
func main() {
	var name string="Cbe"
	const pi float64=3.14412435
	win:=true
	x:=10
	
	fmt.Printf("%.3f  \n", pi) 
	fmt.Printf("%T \n",name)   
	fmt.Printf("%t \n",win)    
	fmt.Printf("%d \n",x)       
	fmt.Printf("%b \n",20)    
	fmt.Printf("%c \n",33)     
	fmt.Printf("%x \n",15)     
	fmt.Printf("%e \n",pi)     
}