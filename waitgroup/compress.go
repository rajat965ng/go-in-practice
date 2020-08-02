package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var i int = -1
	var name string
	for i,name = range os.Args[1:]{
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(name)
	}
	wg.Wait()
	fmt.Printf("Compressed files %d \n",i+1)
}

func compress(name string)  {
	file,err:=os.Open(name)

	if err!=nil {
		panic(err)
	}
	defer file.Close()


	gzfile,err := os.Create(name+".gz")

	if err!=nil{
		panic(err)
	}
	defer gzfile.Close()

	out:=gzip.NewWriter(gzfile)

	io.Copy(out,file)
}