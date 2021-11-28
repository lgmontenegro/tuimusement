package application

import (
	"fmt"
	"lgmontenegro/tuimusement/internal/service"
	"lgmontenegro/tuimusement/internal/task"
)

type App struct {
	//WeatherCrawler service.CrawlerConfig
	CityCrawler service.CrawlerConfig
}

func NewApp(tuiAPIEndpoint string) (appConfig App) {
	fmt.Println("App:", tuiAPIEndpoint)
	cityCrawler := service.NewCrawler(tuiAPIEndpoint)
	return App{
		CityCrawler: cityCrawler,
	}
}

func (c *App) GetCities() (body []byte, err error) {
	body, err = c.CityCrawler.GetData(task.CityFetcher{})
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
