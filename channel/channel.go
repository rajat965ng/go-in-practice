package main

import (
	"fmt"
	"time"
)

func channel(c chan int)  {

	for num:=0;num>=0; {
		num = <-c
		fmt.Println(num, " ")
	}

}

func main() {
	c := make(chan  int)
	a := []int{1,2,3,4,5,0,-1}

    go channel(c)

	for _,i := range a {
		c <- i
	}

	time.Sleep(time.Millisecond*1)

	fmt.Println("End Game")
}
