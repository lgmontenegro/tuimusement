package bootstrap

import (
	"fmt"
	"lgmontenegro/tuimusement/internal/application"
)

func BootstrapApplication(tuiAPIEndpoint string)(app application.App){
	fmt.Println("bootstrap:", tuiAPIEndpoint)
	return application.NewApp(tuiAPIEndpoint)
}