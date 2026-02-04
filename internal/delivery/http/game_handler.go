package http

import (
	"errors"
	"net/http"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/gin-contrib/sessions"
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
	session := sessions.Default(c)

	userID := session.Get("userID")
	var game domain.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Create(c.Request.Context(), &game, userID.(string))
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

func (h *GameHandler) GetAllVerified(c *gin.Context) {
	games, err := h.usecase.GetAllVerified(c.Request.Context())
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
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, game)
}

// Get api/v1/games/my-uploads
func (h *GameHandler) GetByUserId(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")

	games, err := h.usecase.GetByUserId(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}

// PUT api/v1/games/:id
func (h *GameHandler) Update(c *gin.Context) {
	session := sessions.Default(c)
	currentUserID := session.Get("userID").(string)
	currentUserRole := session.Get("role").(string)

	id := c.Param("id")
	var game domain.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Update(c.Request.Context(), id, &game, currentUserID, currentUserRole)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		} else if err.Error() == "permission denied: you are not the owner of this game" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game updated successfully"})
}

// GET /games/:id/reviews
func (h *GameHandler) GetReviewsByGameId(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	reviews, err := h.usecase.GetReviewsByGameId(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
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

// PATCH api/v1/admin/games/:id/verify
func (h *GameHandler) VerifySwitch(c *gin.Context) {
	id := c.Param("id")
	err := h.usecase.VerifySwitch(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game verified/unverified successfully"})
}

func (h *GameHandler) SearchByTitle(c *gin.Context) {
	searchTitle := c.Query("q")
	results, err := h.usecase.SearchByTitle(c.Request.Context(), searchTitle)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search games"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (h *GameHandler) GetUserLibraryWithDetails(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")

	libraryGames, err := h.usecase.GetUserLibraryWithDetails(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, libraryGames)

}

func (h *GameHandler) GetStats(c *gin.Context) {
	result, err := h.usecase.GetStats(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
