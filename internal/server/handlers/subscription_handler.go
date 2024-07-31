package handlers

import (
	"log"
	"net/http"
	"rate/internal/database"
	"rate/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type SubscriptionHandler struct {
	subscriptionRepo database.IEmailRepo
}

func NewSubscriptionHandler(subsRepo database.IEmailRepo) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionRepo: subsRepo,
	}
}

type subscribeInput struct {
	Email string `json:"email"`
}

func (h *SubscriptionHandler) Subscribe(ctx *gin.Context) {
	var req subscribeInput

	if err := ctx.BindJSON(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	err := h.subscriptionRepo.Subscribe(models.Email{
		Email: req.Email,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			ctx.Status(http.StatusConflict)
			return

		}
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)

}
