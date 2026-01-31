package http

import (
	"errors"
	"net/http"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	usecase domain.GameUC
}

func NewGameHandler(usecase domain.GameUC) *GameHandler {
	return &GameHandler{
		usecase: usecase,
	}
}

// POST api/v1/games
func (h *GameHandler) Create(c *gin.Context) {
	var game domain.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Create(c.Request.Context(), &game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Game created successfully"})
}

// GET api/v1/games
func (h *GameHandler) GetAll(c *gin.Context) {
	games, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}

// GET api/v1/games/:id
func (h *GameHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	game, err := h.usecase.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, game)
}

// PUT api/v1/games/:id
func (h *GameHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var game domain.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Update(c.Request.Context(), id, &game)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Game updated successfully"})
}

// DELETE api/v1/games/:id
func (h *GameHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.usecase.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted"})
}
