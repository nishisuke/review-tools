package main

import "sync"

type EmbedsMutex struct {
	key int
	sync.Mutex
}

func main() {
	var v1, v2 int
	println(v1 == v2)
	println(v1 != v2)
	if v1 == 0 && v1 == 0 {
		println("hello, world!")
	}
}
