package webcrawler

import (
	"fmt"
	"net/http"
	"time"
)

// ConnectToWebsite connects to the passed webpage and returns the ioReader
func ConnectToWebsite(webpageURL string) (*http.Response, bool) {
	nilResponse := http.Response{}
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	request, err := http.NewRequest("GET", webpageURL, nil)
	if err != nil {
		fmt.Println("Received error while creating new request: ", err)
		return &nilResponse, false
	}

	request.Header.Set("User-Agent", "GoBot v1.0 https://www.github.com/palvali/GoBot - This bot retrieves links and content.")

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Received error while connecting to website: ", err)
		return &nilResponse, false
	}

	return response, true
}
