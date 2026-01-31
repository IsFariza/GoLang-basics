package http

import (
	"errors"
	"net/http"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/gin-gonic/gin"
)

type EmulationHandler struct {
	usecase domain.EmulationUC
}

func NewEmulationHandler(usecase domain.EmulationUC) *EmulationHandler {
	return &EmulationHandler{
		usecase: usecase,
	}
}

func (h *EmulationHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var emulation domain.Emulation
	if err := c.ShouldBindJSON(&emulation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Create(ctx, &emulation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Emulation created successfully"})
}

func (h *EmulationHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	emulations, err := h.usecase.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emulations)
}

func (h *EmulationHandler) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	task, err := h.usecase.GetById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *EmulationHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	var emulation domain.Emulation

	if err := c.ShouldBindJSON(&emulation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Update(ctx, id, &emulation)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Emulation not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Emulation updated successfully"})
}

func (h *EmulationHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	err := h.usecase.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Emulation not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Emulation deleted"})
}
