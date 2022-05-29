package main

import (
	"sync"
	"time"
)

var (
	money = 100
	lock  = sync.Mutex{}
)

func stingy() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		money += 10
		lock.Unlock()
	}
	println("Stingy Done")
}

func spendy() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		money -= 10
		lock.Unlock()
	}
	println("Spendy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(1000 * time.Millisecond)
	print(money)
}
