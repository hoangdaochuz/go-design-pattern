package singletonpattern

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var lock sync.Mutex
var singletonInstance *Singleton

func GetInstance() *Singleton {
	if singletonInstance == nil {
		// create singleinstacne.
		lock.Lock()
		defer lock.Unlock()
		// We continue to check nil here because when we run a ton of goroutine. In the bad case, it will
		// have 2 or more goroutine by pass previous check nil to be this block code
		// at the same time, only one goroutine run in a lock boundary.
		// if we don't check nil here, the next gorountine will create a single instace again --> invalid Singleton pattern principle
		if singletonInstance == nil {
			fmt.Println("Creating single instance")
			singletonInstance = &Singleton{}
		} else {
			fmt.Println("Instance has already created")
		}
	} else {
		fmt.Println("Instance has already created")
	}
	return singletonInstance
}
