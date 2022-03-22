package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/mars-terminal/openWeatherApi/config"
	"github.com/mars-terminal/openWeatherApi/internal/server/http"
	weatherHandler "github.com/mars-terminal/openWeatherApi/internal/server/http/weather"
	weatherService "github.com/mars-terminal/openWeatherApi/internal/service/weather"
)

var log = logrus.WithField("package", "main")

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	err := config.NewConfig("yaml", "config", "config")
	if err != nil {
		log.WithError(err).Fatal("can't read config")
	}

	var opts = struct {
		OpenWeatherMapApiKey string
	}{
		OpenWeatherMapApiKey: viper.GetString("external.open_weather_map_api_key"),
	}

	log.Info(opts)

	serviceWeather := weatherService.NewService(opts.OpenWeatherMapApiKey)
	handlerWeather := weatherHandler.NewHandler(serviceWeather)

	r := gin.New()

	handlerWeather.InitRoutes(r)

	srv := new(http.Server)
	if err := srv.Run("8000", r); err != nil {
		log.WithError(err).Fatal("error occured while running server server")
	}
}
