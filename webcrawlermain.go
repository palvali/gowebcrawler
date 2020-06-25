package webcrawler

import (
	"sync"
)

// CrawlerMain is the main trigger for webcrawler
func CrawlerMain() {

	sitesChannel := make(chan string)
	crawedLinksChannel := make(chan string)
	pendingCountChannel := make(chan int)

	siteToCrawl := "https://www.crawler-test.com/"

	go func() {
		sitesChannel <- siteToCrawl
		pendingCountChannel <- 1
	}()

	var wg sync.WaitGroup

	go ProcessCrawledLinks(sitesChannel, crawedLinksChannel, pendingCountChannel)
	go MonitorCrawling(sitesChannel, crawedLinksChannel, pendingCountChannel)

	wg.Add(1)
	go CrawlWebpage(&wg, sitesChannel, crawedLinksChannel, pendingCountChannel)
	wg.Wait()
}
