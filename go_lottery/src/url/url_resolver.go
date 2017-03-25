package url

import (
	"net/http"
	"log"
	"io/ioutil"
)

type URLParser struct {
	Url string
}

func (parser URLParser) ReadContent() (string, error) {
	log.Print("Starting ReadContent")

	resp, err := http.Get(parser.Url)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	return string(body), nil

}

func (parser URLParser) generate() string {
	return parser.Url
}
