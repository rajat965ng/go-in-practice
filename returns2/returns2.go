package main

import "fmt"

func names2() (a string, b string)  {
	a = "rajat"
	b = "nigam"
	return
}

func main() {
	n1,n2 := names2()
	fmt.Println(n1)
	fmt.Println(n2)
}