package main

import (
	"log"
	"os"
)

func main() {
	logfile,_ := os.Create("app.log")
	defer logfile.Close()
	logger := log.New(logfile,"app ",log.Lshortfile|log.LstdFlags)
	logger.Println("Before invoking any call")
	logger.Fatal("on getting some error")
}
