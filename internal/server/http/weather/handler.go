package weather

import (
	"github.com/gin-gonic/gin"

	"github.com/mars-terminal/openWeatherApi/internal/service"
)

type Handler struct {
	Service service.WeatherService
}

func NewHandler(service service.WeatherService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes(r gin.IRouter) {
	group := r.Group("/weather")
	group.GET("/:city", h.getWeatherByCityName)
}
