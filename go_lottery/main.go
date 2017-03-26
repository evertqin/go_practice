package main

import (
	"url"
	"log"
)

func main() {
	parser := url.URLParser{"http://broadwayforbrokepeople.com"}

	ch := make(chan string)
	go parser.ParseContent(ch)

	for item := range ch {
		log.Println(item)
	}

}
