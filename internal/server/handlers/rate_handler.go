package handlers

import (
	"net/http"
	"rate/internal/pkg/provider"

	"github.com/gin-gonic/gin"
)

type RateResponse struct {
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

type RateHandler struct {
	rateProvider provider.IRateProvider
}

func NewRateHandler(rateProvider provider.IRateProvider) *RateHandler {
	return &RateHandler{
		rateProvider: rateProvider,
	}
}

func (h *RateHandler) Get(ctx *gin.Context) {

	rate, err := h.rateProvider.GetRate()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch rate"})
	}

	ctx.JSON(http.StatusOK, rate)
}
