package task

import (
	"io/ioutil"
	"lgmontenegro/tuimusement/internal/service"
	"net/http"
)

type CityFetcher struct {
}

func (f CityFetcher) FetchURL(crawlerConfig *service.CrawlerConfig) (body []byte, err error) {

	resp, err := http.Get(crawlerConfig.Endpoint)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
