package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up"
}
