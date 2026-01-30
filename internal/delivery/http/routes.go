package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, gameH *GameHandler) {
	// Base API Group
	api := r.Group("/api/v1")
	{
		// Games API Group
		games := api.Group("/games")
		{
			games.POST("/", gameH.CreateGame)
			games.GET("/", gameH.GetAll)
			games.GET("/:id", gameH.GetById)
			games.PUT("/:id", gameH.Update)
			games.DELETE("/:id", gameH.Delete)
		}
	}
}
