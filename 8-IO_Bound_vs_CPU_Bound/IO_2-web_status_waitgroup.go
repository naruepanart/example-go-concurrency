package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	links := []string{
		"http://hashnode.com",
		"http://dev.to",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://medium.com",
		"http://github.com",
		"http://techcrunch.com",
		"http://techrepublic.com",
	}
	wg.Add(len(links))

	start := time.Now()

	for _, link := range links {
		go checkLink(link)
	}
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func checkLink(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")
		wg.Done()
		return
	}

	fmt.Println(link, "is LIVE!")
	wg.Done()
}
