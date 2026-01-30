package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, gameH *GameHandler, companyH *CompanyHandler) {
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

		companies := api.Group("/companies")
		{
			companies.POST("/", companyH.Create)
			companies.GET("/", companyH.GetAll)
			companies.GET("/:id", companyH.GetById)
			companies.PUT("/:id", companyH.Update)
			companies.DELETE("/:id", companyH.Delete)
		}
	}
}
