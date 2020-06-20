package webcrawler

import (
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func isAnchorTag(tokenType html.TokenType, token html.Token) bool {
	return tokenType == html.StartTagToken && token.DataAtom.String() == "a"
}

func extractLinksFromToken(token html.Token, links *[]string, webpageURL string) {
	for _, attr := range token.Attr {
		if attr.Key == "href" {
			link := attr.Val
			tl := formatURL(webpageURL, getOnlyURLPart(link))

			if !exists(*links, tl) {
				*links = append(*links, tl)
			}
		}
	}
}

func exists(strlist []string, str string) bool {
	for _, s := range strlist {
		if s == str {
			return true
		}
	}
	return false
}

func getOnlyURLPart(l string) string {
	if strings.Contains(l, "#") || strings.Contains(l, "?") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" || strconv.QuoteRune(str) == "'?'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

func formatURL(base string, l string) string {

	base = strings.TrimSuffix(base, "/")

	switch {
	case strings.HasPrefix(l, "https://"):
	case strings.HasPrefix(l, "http://"):
		return l
	case strings.HasPrefix(l, "/"):
		return base + l
	case strings.HasPrefix(l, "#"):
		return base + "/" + l
	default:
		return l
	}
	return l
}
