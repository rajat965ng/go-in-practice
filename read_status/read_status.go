package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn,err := net.Dial("tcp","golang.org:80")
	if err!=nil {
		panic(err.Error())
	}
	fmt.Println(conn,"GET /")

	socket,_ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(socket)
}