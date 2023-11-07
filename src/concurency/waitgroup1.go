package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup

	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("hello world 1")
		time.Sleep(time.Second * time.Duration(rand.Int63n(6)))

	}()

	wait.Add(1)
	go func() {
		defer wait.Done()
		time.Sleep(time.Second * time.Duration(rand.Int63n(6)))
		fmt.Println("hello world 2")
		time.Sleep(time.Second * time.Duration(rand.Int63n(6)))

	}()

	wait.Wait()

	fmt.Println("hello all world")

}
