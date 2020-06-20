package webcrawler

import (
	"strings"

	"golang.org/x/net/html"
)

func isTextTag(tokenType html.TokenType, token html.Token) bool {
	return tokenType == html.TextToken
}

func extractTextFromToken(token html.Token) string {
	data := strings.TrimSpace(token.Data)
	if strings.Contains(data, "function(") || strings.Contains(data, "<iframe") || strings.Contains(data, "<script") {
		return ""
	}
	return data
}
