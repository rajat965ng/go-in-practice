package main

import (
	"fmt"
	"github.com/braintree/manners"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	handler := NewHandler()

	ch := make(chan os.Signal)
	signal.Notify(ch,os.Interrupt,os.Kill)
	go listenOsInterrupt(ch)

	manners.ListenAndServe("0.0.0.0:4000",handler)
}

func listenOsInterrupt(ch chan os.Signal)  {
	<- ch
	manners.Close()
}

func NewHandler() *handler {
	return &handler{}
}

type handler struct {

}

func (h *handler) ServeHTTP(rw http.ResponseWriter, req *http.Request)   {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Printf("Hello %s !!",name)
}