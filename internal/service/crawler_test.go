package service

import (
	"reflect"
	"testing"
)

func TestNewCrawler(t *testing.T) {
	type args struct {
		endpoint string
		key      string
	}
	tests := []struct {
		name    string
		args    args
		wantCfg CrawlerConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCfg := NewCrawler(tt.args.endpoint, tt.args.key); !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("NewCrawler() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}

func TestCrawlerConfig_GetData(t *testing.T) {
	type fields struct {
		Endpoint  string
		Key       string
		Parameter map[string]string
	}
	type args struct {
		crawler Crawler
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantBody []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CrawlerConfig{
				Endpoint:  tt.fields.Endpoint,
				Key:       tt.fields.Key,
				Parameter: tt.fields.Parameter,
			}
			gotBody, err := c.GetData(tt.args.crawler)
			if (err != nil) != tt.wantErr {
				t.Errorf("CrawlerConfig.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("CrawlerConfig.GetData() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}
