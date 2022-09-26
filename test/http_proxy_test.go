package test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestProxy(t *testing.T) {
	proxyUrl, err := url.Parse("http://127.0.0.1:1080")
	if err != nil {
		log.Fatal(err)
	}
	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	//generating the HTTP GET request
	request, err := http.NewRequest("GET", "https://finance.yahoo.com", nil)
	if err != nil {
		log.Println(err)
	}

	//calling the URL
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	//getting the response
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	//printing the response
	s := string(data)
	fmt.Print(s)
}
