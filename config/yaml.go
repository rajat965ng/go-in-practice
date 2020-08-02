package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type yamlPayload struct {
	Enabled bool `json:"enabled"`
	Path string `json:"path"`
}

func main() {
	file,_ :=os.Open("config/config.yml")
	defer file.Close()
	decode := yaml.NewDecoder(file)
	payload := &yamlPayload{}
	decode.Decode(payload)

	fmt.Println(payload)
}

