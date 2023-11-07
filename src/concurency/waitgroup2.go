package main

import (
	"fmt"
	"sync"
)

const NUM int = 5

func main() {
	var wait sync.WaitGroup
	wait.Add(NUM)

	for i := 1; i <= NUM; i++ {
		go print(&wait, i)
	}

	wait.Wait()

	fmt.Println("hello all world")

}

func print(waitGroup *sync.WaitGroup, i int) {
	defer waitGroup.Done()
	fmt.Println("hello ", i)

}
