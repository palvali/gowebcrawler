package webcrawler

import (
	"strings"

	"golang.org/x/net/html"
)

func isAnchorTag(tokenType html.TokenType, token html.Token) bool {
	return tokenType == html.StartTagToken && token.DataAtom.String() == "a"
}

func extractLinksFromToken(token html.Token, webpageURL string) (string, bool) {
	for _, attr := range token.Attr {
		if attr.Key == "href" {
			link := attr.Val
			tl := formatURL(webpageURL, link)
			if tl == "" {
				break
			}
			return tl, true
		}
	}
	return "", false
}

func formatURL(base string, l string) string {

	base = strings.TrimSuffix(base, "/")

	switch {
	case strings.HasPrefix(l, "https://"):
	case strings.HasPrefix(l, "http://"):
		return l
	case strings.HasPrefix(l, "/"):
		return base + l
	}
	return ""
}
