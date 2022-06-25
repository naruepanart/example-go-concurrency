package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} // This creates our waitgroup called, "wg"

func main() {
	wg.Add(2) // This sets our wait group to 2, meaning we will wait for 2 goroutines to complete
	start := time.Now()
	go doSomething()     // Adding the keyword "go" creates a goroutine
	go doSomethingElse() // Adding the keyword "go" creates a goroutine
	wg.Wait()            // This blocks (waits) until wg=0 meaning all goroutines are done.

	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething() {
	fmt.Println("doSomething")
	defer wg.Done() // This decrements wg by one, indicating that doSomething is done.
}

func doSomethingElse() {
	fmt.Println("doSomethingElse")
	defer wg.Done() // This decrements wg by one, indicating that doSomethingElse is done.
}
