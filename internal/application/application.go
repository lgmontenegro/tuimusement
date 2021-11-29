package application

import (
	"encoding/json"
	"fmt"
	"lgmontenegro/tuimusement/internal/domain"
	"lgmontenegro/tuimusement/internal/service"
	"lgmontenegro/tuimusement/internal/task"
	"strings"
)

type App struct {
	WeatherCrawler service.CrawlerConfig
	CityCrawler    service.CrawlerConfig
}

func NewApp(tuiAPIEndpoint, tuiKey, weatherEndpoint, weatherKey string) (appConfig App) {
	var (
		cityCrawler, weatherCrawler service.CrawlerConfig
	)
	cityCrawler = service.NewCrawler(tuiAPIEndpoint, tuiKey)
	weatherCrawler = service.NewCrawler(weatherEndpoint, weatherKey)
	weatherCrawler.Parameter = make(map[string]string)
	weatherCrawler.Parameter["aqi"] = "no"
	weatherCrawler.Parameter["alerts"] = "no"
	weatherCrawler.Parameter["days"] = "2"

	return App{
		CityCrawler:    cityCrawler,
		WeatherCrawler: weatherCrawler,
	}
}

func (c *App) ShowInfo() (err error) {
	cities, err := c.getCities()
	if err != nil {
		return err
	}

	closureGetWeather := func(name string) {
		weather, err := c.getWeather(name)
		if err != nil {
			fmt.Println(err)
		}

		climate := ""
		for _, w := range weather.Forecast.Forecastday {
			climate = fmt.Sprintf("%s - %s", climate, w.Day.Condition.Text)
		}
		climate = strings.TrimLeft(climate, " -")

		_, err = fmt.Printf("Processed city %s | %s \n", name, climate)
		if err != nil {
			fmt.Println(err)
		}
	}

	for _, city := range cities.Cities {
		closureGetWeather(city.Name)
	}

	return nil
}

func (c *App) getCities() (cityCollection domain.CitiesCollection, err error) {
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

func (c *App) getWeather(city string) (weatherStatus domain.Weather, err error) {
	var (
		weather domain.Weather
		body    []byte
	)

	c.WeatherCrawler.Parameter["q"] = city
	body, err = c.WeatherCrawler.GetData(task.WeatherFetcher{})
	if err != nil {
		return domain.Weather{}, err
	}

	json.Unmarshal(body, &weather)

	return weather, nil
}


