package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, gameH *GameHandler, companyH *CompanyHandler, emulationH *EmulationHandler, userH *UserHandler) {
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

		users := api.Group("/users")
		{
			users.POST("/", userH.Create)
			users.GET("/", userH.GetAll)
			users.GET("/:id", userH.GetById)
			users.PUT("/:id", userH.Update)
			users.DELETE("/:id", userH.Delete)
		}

		api.POST("/signup", userH.SignUp)
		api.POST("/login", userH.Login)

		api.GET("/games", gameH.GetAll)
		api.GET("/games/:id", gameH.GetById)

		auth := api.Group("/", AuthMiddleware("user"))
		{
			auth.POST("/games", gameH.Create)
			auth.PUT("/games/:id", gameH.Update)
			auth.DELETE("/games/:id", gameH.Delete)

			auth.POST("/companies", companyH.Create)
			auth.POST("/emulations", emulationH.Create)

			admin := auth.Group("/admin", AuthMiddleware("admin"))
			{
				admin.PATCH("/games/:id/approve", gameH.Approve)

				admin.GET("/users", userH.GetAll)
				admin.DELETE("/users/:id", userH.Delete)
			}
		}
	}
}
