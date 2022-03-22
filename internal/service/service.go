package service

import "github.com/mars-terminal/openWeatherApi/internal/entities/weather"

type WeatherService interface {
	GetWeatherByCityName(string) (*weather.ApiResponse, error)
}
