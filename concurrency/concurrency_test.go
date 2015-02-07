package main

import (
	"testing"
	"sync"
	"log"
)

func Test_UsingWaitGroup(t *testing.T) {
	var done sync.WaitGroup
	for i := 0; i < 4; i++ {
		done.Add(1)
		go func(i int, done *sync.WaitGroup) {
			log.Printf("[-] Doing some stuffs in this %d loop", i)
			done.Done()
		}(i, &done)
	}
	done.Wait()
}
