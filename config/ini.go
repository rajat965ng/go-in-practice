package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
)

func main() {
    config := struct {
		Section struct {
			Enabled bool
			Path string
		}
	}{}

	err := gcfg.ReadFileInto(&config,"config/config.ini")
	if err!=nil {
		panic(err.Error())
	}

	fmt.Println(config)
}