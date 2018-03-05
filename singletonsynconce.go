package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var singInstance *singleton
var inst *singleton
var one sync.Once

func getInstance() *singleton {
	one.Do(func() {
		inst = new(singleton)
	})
	return inst
}

func main() {

	singInstance = getInstance()
	fmt.Println(&singInstance)

	singInstance = getInstance()
	fmt.Println(&singInstance)
}
