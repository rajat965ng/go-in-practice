package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

var opt struct {
	Name string `short:"n" long:"name" default:"world" description:"A name to say hello to."`
	Spanish bool `short:"s" long:"spanish" description:"Use spanish language"`
}

func main() {
	flags.Parse(&opt)

	if opt.Spanish == true {
		fmt.Printf("Hola %s!",opt.Name)
	}else {
		fmt.Printf("Hello %s!",opt.Name)
	}
}