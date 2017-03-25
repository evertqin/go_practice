package main

import (
	"url"
	"fmt"
)

func main() {
	parser := url.URLParser{"http://broadwayforbrokepeople.com"}

	body, _ := parser.ReadContent()

	fmt.Println(body)


}
