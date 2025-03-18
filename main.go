package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex
var hashmap = make(map[int]bool, 0)

func printOnce(num int) {
	m.Lock()
	defer m.Unlock()
	if _, exist := hashmap[num]; !exist {
		fmt.Println(num)
	}
	hashmap[num] = true
}

func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(1 * time.Second)
}
