package webcrawler

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Extract retrieves the information from the webpage body
func extractContent(httpBody io.Reader, webpageURL string) ([]string, map[string]int) {

	links := []string{}
	wordscount := map[string]int{}

	page := html.NewTokenizer(httpBody)

	for {
		tokenType := page.Next()

		if tokenType == html.ErrorToken {
			return links, wordscount
		}

		token := page.Token()

		if isAnchorTag(tokenType, token) {
			extractLinksFromToken(token, &links, webpageURL)
		} else if isTextTag(tokenType, token) {
			extractedwords := extractTextFromToken(token)
			addWordsToMap(extractedwords, wordscount)
		}
	}
}

func addWordsToMap(data string, wordscount map[string]int) {
	if len(data) == 0 {
		return
	}
	words := strings.Fields(data)
	for _, word := range words {
		wordscount[word]++
	}
}