package handlers

import (
	"esapp/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) GetPaymentHistory(c *gin.Context) {

	val, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := val.(uint)

	history, err := h.paymentService.GetPaymentHistory(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch payment history",
		})
		return
	}

	c.JSON(http.StatusOK, history)
}
