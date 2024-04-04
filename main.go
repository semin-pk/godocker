package main
import (
	"fmt"
	"time"
)

func main(){
	for {
		fmt.Println("Hello Go")
		time.Sleep(10*time.Second)
	}
}