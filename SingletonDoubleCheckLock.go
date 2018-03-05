package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var singInstance *singleton

var mt = &sync.RWMutex{}
var inst *singleton

func getInstance() *singleton {
	mt.RLock()
	if inst == nil {
		mt.RUnlock()
		mt.Lock()
		defer mt.Unlock()
		if inst == nil {
			inst = new(singleton)
		}
		return inst
	} else {
		mt.RUnlock()
		return inst
	}
}

func main() {

	singInstance = getInstance()
	fmt.Println(&singInstance)

	singInstance = getInstance()
	fmt.Println(&singInstance)
}
