package main
import (
	"fmt"
	"errors"
)

func main() {
			total,error := (sum(12,10))
			if error!= nil{
				fmt.Println("Error:",error)
			} else {
				fmt.Println(total)
			}

}

func sum(start int,end int) (int, error) {
	if start > end
	{
		return 0, errors.New("start point is greater than end point")
	}
	total := 0
	for i:= start; i<=end; i++{
		total = total +i
	}
	return total,nil
}

}