package task

import (
	"lgmontenegro/tuimusement/internal/service"
	"reflect"
	"testing"
)

func TestWeatherFetcher_FetchURL(t *testing.T) {
	type args struct {
		crawlerConfig *service.CrawlerConfig
	}
	tests := []struct {
		name     string
		f        WeatherFetcher
		args     args
		wantBody []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := WeatherFetcher{}
			gotBody, err := f.FetchURL(tt.args.crawlerConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("WeatherFetcher.FetchURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("WeatherFetcher.FetchURL() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}
