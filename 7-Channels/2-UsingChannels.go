package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func main() {

	start := time.Now()
	go doSomething()
	go doSomethingElse()

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething() {
	ch <- "doSomething"
}

func doSomethingElse() {
	ch <- "doSomethingElse"
}
