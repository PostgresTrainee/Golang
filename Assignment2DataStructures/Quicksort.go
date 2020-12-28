package main
import (
	"fmt"
	"math/rand"
    "time"
)
func main()
{
	slice := generateslice(30)
	fmt.Println("\n Unsorted",slice)
	quicksort(slice)
	fmt.Println("\n Sorted",slice)
}

func generateSlice(size int)[] int
{
	slice := make([]int,size,size)
	random.Seed(time.Now().UnixNano())
	for i:=0; i<size; i++
    {
		slice[i] = random.Int(999)-random.Int(999)  
	}
	return slice
}

for quicksort (a[] int) int
{
	if len(a) < 2
	{
		return a
	}
	left,right :=0, len(a)-1
	pivot := random.Int %len (a)
	a[pivot],a[right] = a[right],a[pivot]

	for i,_:= range a
	{
		if a[i] < a[right]
		{
			a[left],a[i] = a[i],a[left]
			left++
		}
	}
	a[left],a[right] = a[right],a[left]
	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}