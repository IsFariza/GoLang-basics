package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, gameH *GameHandler, companyH *CompanyHandler, emulationH *EmulationHandler, purchaseH *PurchaseHandler, reviewH *ReviewHandler) {
	// Base API Group
	api := r.Group("/api/v1")
	{
		// Games API Group
		games := api.Group("/games")
		{
			games.POST("/", gameH.Create)
			games.GET("/", gameH.GetAll)
			games.GET("/:id", gameH.GetById)
			games.PUT("/:id", gameH.Update)
			games.DELETE("/:id", gameH.Delete)
			games.GET("/:id/reviews", gameH.GetReviewsByGameId)
		}

		companies := api.Group("/companies")
		{
			companies.POST("/", companyH.Create)
			companies.GET("/", companyH.GetAll)
			companies.GET("/:id", companyH.GetById)
			companies.PUT("/:id", companyH.Update)
			companies.DELETE("/:id", companyH.Delete)
		}

		emulations := api.Group("/emulations")
		{
			emulations.POST("/", emulationH.Create)
			emulations.GET("/", emulationH.GetAll)
			emulations.GET("/:id", emulationH.GetById)
			emulations.PUT("/:id", emulationH.Update)
			emulations.DELETE("/:id", emulationH.Delete)
		}
		purchases := api.Group("/purchases")
		{
			purchases.POST("/", purchaseH.Create)
			purchases.GET("/", purchaseH.GetAll)
			purchases.GET("/:id", purchaseH.GetById)
			purchases.DELETE("/:id", purchaseH.Delete)
		}
		reviews := api.Group("/reviews")
		{
			reviews.POST("/", reviewH.Create)
			reviews.GET("/", reviewH.GetAll)
			reviews.GET("/:id", reviewH.GetById)
			reviews.PUT("/:id", reviewH.Update)
			reviews.DELETE("/:id", reviewH.Delete)
		}
	}
}
