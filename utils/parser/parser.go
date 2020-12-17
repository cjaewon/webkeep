package parser

import (
	"io/ioutil"
	"net/http"
	"time"
)

// Fetch fetchs html text from url
func Fetch(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)

	return str, nil

}
