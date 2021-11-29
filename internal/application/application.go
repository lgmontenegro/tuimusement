package application

import (
	"encoding/json"
	"fmt"
	"lgmontenegro/tuimusement/internal/domain"
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

func (c *App) GetCities() (cityCollection domain.CitiesCollection, err error) {
	var (
		cities domain.CitiesCollection		
		body   []byte
	)

	body, err = c.CityCrawler.GetData(task.CityFetcher{})
	if err != nil {
		return domain.CitiesCollection{}, err
	}

	json.Unmarshal(body, &cities.Cities)

	return cities, nil
}
