package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cc := &http.Client{Timeout:time.Second}
	res,err := cc.Get("http://goinpracticebook.com")
	if err!=nil {
		panic(err)
	}
	b,_ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s",b)
}
