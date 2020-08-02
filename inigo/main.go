package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprint(res,"Hello world this is sample response")
}

func main() {

	http.HandleFunc("/hello",hello)
	http.ListenAndServe("0.0.0.0:4000",nil)

}
