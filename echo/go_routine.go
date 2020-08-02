package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside go routine !!")
	go func() {
		fmt.Println("inside go routine !!")
	}()
	fmt.Println("outside again !!")

	runtime.Gosched()
}
