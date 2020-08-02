package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	state := make(map[int]int)
    var mutex = &sync.Mutex{}
	var readOps uint64
	var writeOps uint64

	go func() {
		total:=0
		for i:=0;i<100;i++ {
			key := rand.Intn(5)
			mutex.Lock()
			total+=state[key]
			mutex.Unlock()
			atomic.AddUint64(&readOps,1)
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		for i:=0;i<100;i++ {
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			atomic.AddUint64(&writeOps,1)
			time.Sleep(time.Millisecond)
		}
	}()

    time.Sleep(time.Second)
    fmt.Printf("ReadOps %d \n",atomic.LoadUint64(&readOps))
	fmt.Printf("WriteOps %d \n",atomic.LoadUint64(&writeOps))

    mutex.Lock()
    fmt.Printf("The state is {}",state)
    mutex.Unlock()
}
