package main

import "fmt"

type Singleton struct {
}

var singInstance *Singleton
var inst *Singleton

func getInstance() *Singleton {
	if inst == nil {
		inst = new(Singleton)
	}
	return inst
}

func main() {

	singInstance = getInstance()
	fmt.Println(&singInstance)

	singInstance = getInstance()
	fmt.Println(&singInstance)
}
