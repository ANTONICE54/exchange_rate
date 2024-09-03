package handlers

import (
	"net/http"
	"rate/internal/apperrors"
	"rate/internal/models"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=subscription_handler.go -destination=mock/subscription_handler.go

type ISubscriptionService interface {
	Subscribe(email models.Email) (*models.Email, error)
}

type SubscriptionHandler struct {
	subscriptionService ISubscriptionService
}

func NewSubscriptionHandler(subsService ISubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionService: subsService,
	}
}

type subscribeInput struct {
	Email string `json:"email"`
}

func (h *SubscriptionHandler) Subscribe(ctx *gin.Context) {
	var req subscribeInput

	if err := ctx.BindJSON(&req); err != nil {
		err = apperrors.ErrBadRequest
		ctx.JSON(apperrors.Status(err), gin.H{"error": err.Error()})
		return
	}

	_, err := h.subscriptionService.Subscribe(models.Email{
		Email: req.Email,
	})

	if err != nil {

		ctx.JSON(apperrors.Status(err), gin.H{"error": err.Error()})

		return
	}

	ctx.Status(http.StatusOK)

}
