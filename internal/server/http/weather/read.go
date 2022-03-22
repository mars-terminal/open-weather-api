package weather

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getWeatherByCityName(ctx *gin.Context) {
	city := ctx.Param("city")
	if len(city) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "city name length is zero",
		})
		return
	}

	apiResponse, err := h.Service.GetWeatherByCityName(city)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apiResponse)
}
