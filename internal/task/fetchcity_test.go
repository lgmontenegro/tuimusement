package task

import (
	"lgmontenegro/tuimusement/internal/service"
	"reflect"
	"testing"
)

func TestCityFetcher_FetchURL(t *testing.T) {
	type args struct {
		crawlerConfig *service.CrawlerConfig
	}
	tests := []struct {
		name     string
		f        CityFetcher
		args     args
		wantBody []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := CityFetcher{}
			gotBody, err := f.FetchURL(tt.args.crawlerConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("CityFetcher.FetchURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("CityFetcher.FetchURL() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}
