package main

import (
	"fmt"
)

var result = 0
var value = 97

func main() {
	goChan := make(chan int)
	/* mainChan := make(chan string) */
	go calculateSquare(value, goChan)
	/* go reportResult(goChan, mainChan) */
	fmt.Println("The result of", value, "squared", "is", <-goChan)

}
func calculateSquare(value int, goChan chan int) {
	result = value * value
	goChan <- result

}
