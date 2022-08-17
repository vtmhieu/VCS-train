package main

import (
	"fmt"
	"net/http"
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

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		c <- "might be down!"
		//send a message to channel
		return
	}

	fmt.Println(link, "is up!")
	c <- "is up!"
}
