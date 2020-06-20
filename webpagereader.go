package webcrawler

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// CrawlWebpage crawls a page and returns the site in string
func CrawlWebpage(webpageURL string) io.Reader {
	client := http.Client{
		Timeout: 20 * time.Second,
	}

	request, err := http.NewRequest("GET", webpageURL, nil)
	if err != nil {
		fmt.Println("Received error while creating new request: ", err)
		return nil
	}

	request.Header.Set("User-Agent", "GoBot v1.0 https://www.github.com/palvali/GoBot - This bot retrieves links and content.")

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Received error while connecting to website: ", err)
		return nil
	}

	respBody := response.Body
	defer respBody.Close()

	links, wordscount := extractContent(respBody, webpageURL)
	fmt.Println(links)
	fmt.Println(wordscount)

	return respBody
}
