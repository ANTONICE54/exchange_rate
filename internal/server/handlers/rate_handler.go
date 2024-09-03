package handlers

import (
	"net/http"
	"rate/internal/apperrors"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=rate_handler.go -destination=mock/rate_handler.go

type RateResponse struct {
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

type IRateProviderService interface {
	Get() (*float64, error)
}

type RateHandler struct {
	rateProviderService IRateProviderService
}

func NewRateHandler(RateProviderService IRateProviderService) *RateHandler {
	return &RateHandler{
		rateProviderService: RateProviderService,
	}
}

func (h *RateHandler) Get(ctx *gin.Context) {

	rate, err := h.rateProviderService.Get()

	if err != nil {

		ctx.JSON(apperrors.Status(err), gin.H{"error": err.Error()})

		return

	}

	ctx.JSON(http.StatusOK, rate)
}
