package bootstrap

import (
	"lgmontenegro/tuimusement/internal/application"
)

func BootstrapApplication(tuiAPIEndpoint, tuiKey, weatherEndpoint, weatherKey string) (app application.App) {
	return application.NewApp(tuiAPIEndpoint, tuiKey, weatherEndpoint, weatherKey)
}
