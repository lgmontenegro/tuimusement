package service

import "fmt"

type Crawler interface {
	FetchURL(crawlerConfig *CrawlerConfig) (body []byte, err error)
}

type CrawlerConfig struct {
	Endpoint string
	Key      string
	Parameter map[string]string
}

func NewCrawler(endpoint, key string) (cfg CrawlerConfig) {
	fmt.Println("crawler:", endpoint)
	return CrawlerConfig {
		Endpoint: endpoint,
		Key:      key,
	}
}

func (c *CrawlerConfig) GetData(crawler Crawler) (body []byte, err error) {
	return crawler.FetchURL(c)
}
