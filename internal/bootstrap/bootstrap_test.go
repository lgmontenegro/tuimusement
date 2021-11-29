package bootstrap

import (
	"lgmontenegro/tuimusement/internal/application"
	"reflect"
	"testing"
)

func TestBootstrapApplication(t *testing.T) {
	type args struct {
		tuiAPIEndpoint  string
		tuiKey          string
		weatherEndpoint string
		weatherKey      string
	}
	tests := []struct {
		name    string
		args    args
		wantApp application.App
	}{
		{
			name: "normal bootstraping",
			args: args {
				tuiAPIEndpoint: "",
				tuiKey: "",
				weatherEndpoint: "",
				weatherKey: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotApp := BootstrapApplication(tt.args.tuiAPIEndpoint, tt.args.tuiKey, tt.args.weatherEndpoint, tt.args.weatherKey); !reflect.DeepEqual(gotApp, tt.wantApp) {
				t.Errorf("BootstrapApplication() = %v, want %v", gotApp, tt.wantApp)
			}
		})
	}
}
