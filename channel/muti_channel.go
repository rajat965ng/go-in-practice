package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done:=time.After(30*time.Second)
	echo:=make(chan []byte)

	go read(echo)
	for   {
		select {
		case data:=<-echo:
			os.Stdout.Write(data)
		case <-done:
			fmt.Println("Time over !!")
			os.Exit(0)
		}
	}


}

func read(out chan []byte)  {
	for   {
		data := make([]byte,1024)
		n,_ := os.Stdin.Read(data)

		if n>0 {
			out <- data
		}
	}
}