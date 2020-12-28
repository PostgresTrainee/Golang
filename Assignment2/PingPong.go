package main
import (
	"fmt"
	"time"
)

func ping(awake chan string)
{
	for i:=1; i<6; i++
	{
		awake <- "ping"
	}
}

func sleep(awake chan string){
for {
	   msg := <- click
	   fmt.Println(msg)
	   time.Sleep(time.Second * 1)   
}
}

func pong(awake chan string){
for i=1;i<6;i++
{
  awake <- "ping"            
}
}

func main()
{
	var awake chan string = make(chan string)
	go ping(awake)
	go sleep(awake)
	go pong(awake)

	var input string
	fmt.Scanln(&input)
}


}