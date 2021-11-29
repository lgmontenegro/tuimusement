package task

import (
	"io/ioutil"
	"lgmontenegro/tuimusement/internal/service"
	"net/http"
	"net/url"
	"time"
)

type WeatherFetcher struct {
}

func (f WeatherFetcher) FetchURL(crawlerConfig *service.CrawlerConfig) (body []byte, err error) {
	baseUrl, err := url.Parse(crawlerConfig.Endpoint)
	if err != nil {
		return []byte{}, err
	}

	param := url.Values{}
	param.Add("key", crawlerConfig.Key)
	for queryString, value := range crawlerConfig.Parameter {
		param.Add(queryString, value)
	}
	baseUrl.RawQuery = param.Encode()

	client := http.Client {
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(baseUrl.String())
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
