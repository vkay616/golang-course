package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.amazon.in",
		"https://www.stackoverflow.com",
		"https://www.facebook.com",
		"https://go.dev",
		"https://github.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		// go checkLink(l, c)
		new_link := l
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(new_link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up and running!")
	c <- link
}
