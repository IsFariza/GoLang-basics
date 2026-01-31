package http

import (
	"errors"
	"net/http"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	usecase domain.CompanyUC
}

func NewCompanyHandler(usecase domain.CompanyUC) *CompanyHandler {
	return &CompanyHandler{
		usecase: usecase,
	}
}

func (h *CompanyHandler) Create(c *gin.Context) {
	var company domain.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	h.usecase.Create(ctx, &company)

	c.JSON(http.StatusCreated, gin.H{"message": "Company created successfully"})
}

func (h *CompanyHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	companies, err := h.usecase.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, companies)
}

func (h *CompanyHandler) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	companies, err := h.usecase.GetById(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, companies)
}

func (h *CompanyHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var company domain.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Update(ctx, id, &company)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated"})
}

func (h *CompanyHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	err := h.usecase.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}
