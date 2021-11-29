package bootstrap

import (
	"lgmontenegro/tuimusement/internal/application"
)

func BootstrapApplication(tuiAPIEndpoint, weatherEndpoint, weatherKey string) (app application.App) {
	return application.NewApp(tuiAPIEndpoint, weatherEndpoint, weatherKey)
}
