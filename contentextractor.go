package webcrawler

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

// Extract retrieves the information from the webpage body
func extractContent(webpageURL string, crawedLinksChannel chan string) {

	client := http.Client{
		Timeout: 60 * time.Second,
	}

	request, err := http.NewRequest("GET", webpageURL, nil)
	if err != nil {
		fmt.Println("Received error while creating new request: ", err)
		return
	}

	request.Header.Set("User-Agent", "GoBot v1.0 https://www.github.com/palvali/GoBot - This bot retrieves links and content.")

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Received error while connecting to website: ", err)
		return
	}

	defer response.Body.Close()

	tokenizer := html.NewTokenizer(response.Body)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			return
		}

		token := tokenizer.Token()

		if isAnchorTag(tokenType, token) {
			cl, ok := extractLinksFromToken(token, webpageURL)

			if ok {
				go func() {
					crawedLinksChannel <- cl
				}()
			}
		}
	}
}
