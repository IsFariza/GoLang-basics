package http

import (
	"errors"
	"net/http"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/BlackHole55/software-store-final/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	usecase domain.ReviewUC
}

func NewReviewHandler(usecase *usecase.ReviewUsecase) *ReviewHandler {
	return &ReviewHandler{usecase: usecase}
}

func (h *ReviewHandler) Create(c *gin.Context) {
	var review domain.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Create(c.Request.Context(), &review); err != nil {
		if err.Error() == "review denied: you must own the game to leave a review" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Review submitted successfully"})
}

func (h *ReviewHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	reviews, err := h.usecase.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, reviews)
}
func (h *ReviewHandler) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	review, err := h.usecase.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	userIdInterface, exists := c.Get("currentUserID")
	if !exists || userIdInterface == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Please login"})
		return
	}
	currentUserID, ok := userIdInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user context"})
		return
	}
	var input struct {
		Content string `json:"content" binding:"required"`
		Rating  int    `json:"rating" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updates := &domain.Review{
		Content: input.Content,
		Rating:  &input.Rating,
	}

	if err := h.usecase.Update(ctx, id, currentUserID, updates); err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}
func (h *ReviewHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	userIdInterface, existsId := c.Get("currentUserID")
	userRoleInterface, existsRole := c.Get("currentUserRole")

	if !existsId || !existsRole || userIdInterface == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User session not found"})
		return
	}

	userId, ok1 := userIdInterface.(string)
	userRole, ok2 := userRoleInterface.(string)

	if !ok1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format"})
		return
	}
	if !ok2 {
		userRole = "user"
	}

	ctx := c.Request.Context()

	if err := h.usecase.Delete(ctx, id, userId, userRole); err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
