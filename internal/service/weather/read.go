package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mars-terminal/openWeatherApi/internal/entities/weather"
)

const uri = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"

func (s *Service) GetWeatherByCityName(city string) (*weather.ApiResponse, error) {
	logger := log.WithField("handler", "GetWeatherByCityName")

	response, err := http.Get(fmt.Sprintf(uri, city, s.ApiKey))
	if err != nil {
		logger.WithError(err).Error("can't make request to call open weather api")
		return nil, err
	}
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.WithError(err).Error("can't read request body after call open weather api")
		return nil, err
	}

	resp := &weather.ApiResponse{}

	if err := resp.Unmarshal(all); err != nil {
		logger.WithError(err).Error("can't unmarshal request body after read open weather api data")
		return nil, err
	}

	if resp.Message != "" {
		logger.Error(resp.Message)
		return nil, fmt.Errorf(resp.Message)
	}

	return resp, nil
}
