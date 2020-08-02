package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go copy(os.Stdin,os.Stdout)
	time.Sleep(30*time.Second)
	fmt.Println("Time over !!")
	os.Exit(0)
}

func copy(reader io.Reader,writer io.Writer)  {
	io.Copy(writer,reader)
}