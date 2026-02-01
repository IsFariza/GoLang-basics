package http

import (
	"errors"
	"net/http"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
	usecase domain.PurchaseUC
}

func NewPurchaseHandler(usecase domain.PurchaseUC) *PurchaseHandler {
	return &PurchaseHandler{usecase: usecase}
}
func (h *PurchaseHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var purchase domain.Purchase
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Create(ctx, &purchase); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Purchase created successfully"})
}

func (h *PurchaseHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	purchases, err := h.usecase.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, purchases)
}
func (h *PurchaseHandler) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	purchase, err := h.usecase.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Purchase not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, purchase)
}

func (h *PurchaseHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if err := h.usecase.Delete(ctx, id); err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Purchase not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Purchase deleted successfully"})
}
