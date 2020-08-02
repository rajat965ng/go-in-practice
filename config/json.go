package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type payload struct {
	Enabled bool `json:"enabled"`
	Path string `json:"path"`
}

func main() {
	data,err := ioutil.ReadFile("config/config.json")
	if err!=nil {
		panic(err.Error())
	}
	payload := &payload{}
	json.Unmarshal(data,payload)

	fmt.Println(payload)

}
