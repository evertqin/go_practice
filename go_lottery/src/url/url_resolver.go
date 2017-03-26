package url

import (
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"golang.org/x/net/html"
)

const (
	tempFilename = "go_lottery/tmp/output.html"
)

type URLParser struct {
	Url string
}

func check(err error) {
	if err != nil {
		log.Fatalf("Something is wrong: %s", err)
		panic(err)
	}
}

func (parser URLParser) readContent() ([]byte, error) {
	log.Print("Starting ReadContent")

	resp, err := http.Get(parser.Url)

	check(err)

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	check(err)

	return body, nil
}

func (parser URLParser) writeToFile(input []byte) {
	log.Print("WriteToFile")

	err := ioutil.WriteFile(tempFilename, input, 0644)

	if err != nil {
		log.Fatalf("Some errors: %s", err)
		panic(err)
	}

}

func (parser URLParser) Generate() []byte {
	var content []byte
	if _, err := os.Stat(tempFilename); os.IsNotExist(err) {
		content, err = parser.readContent()
		check(err)
	} else {
		content, err = ioutil.ReadFile(tempFilename)
		check(err)
	}

	return content

}

func (parser URLParser) ParseContent(outUrl chan<- string) {

	resp, err := http.Get(parser.Url)

	check(err)

	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			close(outUrl)
			return
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				log.Printf("We found a link: %s", t.Data)

				for _, a := range t.Attr {
					outUrl <- a.Val
				}
			}

		}

	}
}
