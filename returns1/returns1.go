package main

import "fmt"

func Names() (string,string) {
	return "rajat", "nigam"
}

func main() {
	a,b:= Names()
	fmt.Println(a)
	fmt.Println(b)
}