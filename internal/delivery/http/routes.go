package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, gameH *GameHandler, companyH *CompanyHandler, emulationH *EmulationHandler) {
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
		}

		companies := api.Group("/companies")
		{
			companies.POST("/", companyH.CreateCompany)
			companies.GET("/", companyH.GetAll)
			companies.GET("/:id", companyH.GetById)
			companies.PUT("/:id", companyH.Update)
			companies.DELETE("/:id", companyH.Delete)
		}

		emulations := api.Group("/emulations")
		{
			emulations.POST("/", emulationH.CreateEmulation)
			emulations.GET("/", emulationH.GetAll)
			emulations.GET("/:id", emulationH.GetById)
			emulations.PUT("/:id", emulationH.Update)
			emulations.DELETE("/:id", emulationH.Delete)
		}
	}
}
