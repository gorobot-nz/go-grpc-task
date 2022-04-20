package bitlyClient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type BitlyClient struct {
	token  string
	client *http.Client
}

func NewBitlyClient(token string) *BitlyClient {
	return &BitlyClient{token: token, client: &http.Client{}}
}

func (c *BitlyClient) ShortLink(url string) (string, error) {
	data := strings.NewReader(fmt.Sprintf(`{ "long_url": %s }`, url))
	req, err := http.NewRequest("POST", "https://api-ssl.bitly.com/v4/shorten", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer {TOKEN}")
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	bodyText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyText), nil
}