package webcrawler

import (
	"fmt"
	"sync"
)

// CrawlWebpage crawls a page
func CrawlWebpage(wg *sync.WaitGroup, sitesChannel chan string, crawedLinksChannel chan string, pendingCountChannel chan int) {

	crawledSites := 0

	for webpageURL := range sitesChannel {
		extractContent(webpageURL, crawedLinksChannel)
		pendingCountChannel <- -1
		crawledSites++
	}

	fmt.Println("Crawled ", crawledSites, " web pages.")

	wg.Done()
}
