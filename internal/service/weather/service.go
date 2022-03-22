package weather

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(
	logrus.Fields{
		"package": "weather",
		"layer":   "service",
	},
)

type Service struct {
	ApiKey string
}

func NewService(apiKey string) *Service {
	return &Service{ApiKey: apiKey}
}
