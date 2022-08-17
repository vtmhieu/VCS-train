package main

import (
	"fmt"
	"net/http"
	"time"
)

// go routine = main routine + child go routine
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	c := make(chan string)
	//makeing a new channel type string

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		c <- link
		//send a message to channel
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
