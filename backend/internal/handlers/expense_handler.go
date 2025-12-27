package handlers

import (
	"esapp/internal/dto"
	"esapp/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	expenseService services.ExpenseService
}

func NewExpenseHandler(expenseService services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{expenseService: expenseService}
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {

	val, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := val.(uint)

	var req dto.CreateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.expenseService.CreateExpense(userID, req); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{
		"message": "expense created successfully",
	})
}

func (h *ExpenseHandler) ListExpenses(c *gin.Context) {

	val, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := val.(uint)

	groupIDParam := c.Param("groupId")
	groupID, err := strconv.ParseUint(groupIDParam, 10, 64)
	if err != nil || groupID == 0 {
		c.JSON(400, gin.H{"error": "invalid group id"})
		return
	}

	expenses, err := h.expenseService.ListExpenses(userID, uint(groupID))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	}

	response := make([]gin.H, 0)
	for _, e := range expenses {
		response = append(response, gin.H{
			"id":         e.ID,
			"title":      e.Title,
			"amount":     e.Amount,
			"paid_by_id": e.PaidByID,
			"created_at": e.CreatedAt,
		})
	}

	c.JSON(200, response)
}

func (h *ExpenseHandler) CalculateBalances(c *gin.Context) {

	val, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := val.(uint)
	groupIDParam := c.Param("groupId")
	groupID, err := strconv.ParseUint(groupIDParam, 10, 64)
	if err != nil || groupID == 0 {
		c.JSON(400, gin.H{"error": "invalid group id"})
		return
	}

	balances, err := h.expenseService.CalculateBalances(userID, uint(groupID))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, balances)
}
