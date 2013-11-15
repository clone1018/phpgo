package phpgo

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func FileGetContents(filename string) (string, error) {
	var output string

	url, err := url.Parse(filename)
	if err != nil {
		log.Panic(err)
	}

	if url.Scheme == "https" || url.Scheme == "http" {
		// http request
		resp, err := http.Get(filename)
		defer resp.Body.Close()
		if err != nil {
			return "", err
		}

		bytes, _ := ioutil.ReadAll(resp.Body)
		output = string(bytes[:])

	} else {
		// fs request
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			// Failure to read, abort!
			return "", err
		}

		output = string(bytes[:])
	}

	return output, nil
}
