package webcrawler

import (
	"fmt"
	"sync"
)

// CrawlWebpage crawls a page
func CrawlWebpage(wg *sync.WaitGroup, sitesChannel chan string, crawedLinksChannel chan string, pendingCountChannel chan int) {

	crawledSites := 0

	for webpageURL := range sitesChannel {
		fmt.Println("Crawling site: ", webpageURL)

		extractContent(webpageURL, crawedLinksChannel)
		crawledSites++

		go func() {
			pendingCountChannel <- -1
		}()
	}

	fmt.Println("Crawled ", crawledSites, " web pages.")
	close(crawedLinksChannel)
	close(pendingCountChannel)

	wg.Done()
}
