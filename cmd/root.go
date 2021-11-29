package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"lgmontenegro/tuimusement/internal/application"
	"lgmontenegro/tuimusement/internal/bootstrap"
)

var (
	app     application.App
	rootCmd = &cobra.Command{
		Use:   "tuimusement",
		Short: "Tui Musement CLI will show you all the weather of all the available Musements city",
		RunE: func(cmd *cobra.Command, args []string) error {
			app = bootstrap.BootstrapApplication(viper.GetString("tui_musement_api"))
			cities, err := app.GetCities()
			if err != nil {
				return err
			}

			for i,city := range cities.Cities {
				fmt.Println(i, city.Name)
			}

			return nil
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	configFilePath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(configFilePath)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configFilePath)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
