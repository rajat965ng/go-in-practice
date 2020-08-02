package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp,_ := http.Get("http://example.com/")
	buff,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(buff))
	resp.Body.Close()
}
