package main
import (
	"fmt"
	"time"
)
func main() {
	s := make(chan int, 3)
	w := 10
	go worker(s, 1)
	go worker(s, 2)
	go worker(s, 3)
	go generateInts(s, w)
	time.Sleep(10 * time.Second)
}
func worker(select {
case condition:
	
} chan int, id int) {
	for i := range s {
		fmt.Printf("Worker %d received value %d\n", id, i)
		time.Sleep(1 * time.Second)
	}
}
func generateInts(c chan int, q int) {
	for i := 0; i < q; i++ {
		c <- i
		fmt.Printf("Sent value %d\n", i)
		time.Sleep(1 * time.Millisecond)
	}
	close(s)
}
