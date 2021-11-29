package application

import (
	"lgmontenegro/tuimusement/internal/service"
	"reflect"
	"testing"
)

func TestNewApp(t *testing.T) {
	type args struct {
		tuiAPIEndpoint  string
		weatherEndpoint string
		weatherKey      string
	}
	tests := []struct {
		name          string
		args          args
		wantAppConfig App
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAppConfig := NewApp(tt.args.tuiAPIEndpoint, tt.args.weatherEndpoint, tt.args.weatherKey); !reflect.DeepEqual(gotAppConfig, tt.wantAppConfig) {
				t.Errorf("NewApp() = %v, want %v", gotAppConfig, tt.wantAppConfig)
			}
		})
	}
}

func TestApp_ShowInfo(t *testing.T) {
	type fields struct {
		WeatherCrawler service.CrawlerConfig
		CityCrawler    service.CrawlerConfig
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &App{
				WeatherCrawler: tt.fields.WeatherCrawler,
				CityCrawler:    tt.fields.CityCrawler,
			}
			if err := c.ShowInfo(); (err != nil) != tt.wantErr {
				t.Errorf("App.ShowInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
