package webcrawler

import (
	"fmt"

	"golang.org/x/net/html"
)

// Extract retrieves the information from the webpage body
func extractContent(webpageURL string, crawedLinksChannel chan string) {

	response, success := ConnectToWebsite(webpageURL)

	if !success {
		fmt.Println("Received error while connecting to website: ", webpageURL)
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
