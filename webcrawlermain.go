package webcrawler

import (
	"sync"
)

// CrawlerMain is the main trigger for webcrawler
func CrawlerMain() {

	sitesChannel := make(chan string)
	crawedLinksChannel := make(chan string)
	pendingCountChannel := make(chan int)

	siteToCrawl := "https://theuselessweb.com/"

	go func() {
		crawedLinksChannel <- siteToCrawl
	}()

	var wg sync.WaitGroup

	go ProcessCrawledLinks(sitesChannel, crawedLinksChannel, pendingCountChannel)
	go MonitorCrawling(sitesChannel, crawedLinksChannel, pendingCountChannel)

	var numCrawlerThreads = 50
	for i := 0; i < numCrawlerThreads; i++ {
		wg.Add(1)
		go CrawlWebpage(&wg, sitesChannel, crawedLinksChannel, pendingCountChannel)
	}

	wg.Wait()
}
