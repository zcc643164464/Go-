package main

import (
	"fmt"
	"sync"
)

func main() {

	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		defer memoryAccess.Unlock()
		value++
	}()
	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n",value)
	}else{
		fmt.Printf("the value is %v.\n",value)
	}
	memoryAccess.Unlock()

}