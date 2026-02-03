package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, gameH *GameHandler, companyH *CompanyHandler, emulationH *EmulationHandler, userH *UserHandler, purchaseH *PurchaseHandler, reviewH *ReviewHandler) {
	api := r.Group("/api/v1")
	{
		api.POST("/signup", userH.SignUp)
		api.POST("/login", userH.Login)

		api.GET("/games", gameH.GetAll)
		api.GET("/games/:id", gameH.GetById)
		api.GET("/games/:id/reviews", gameH.GetReviewsByGameId)
		api.GET("/companies", companyH.GetAll)
		api.GET("/companies/:id", companyH.GetById)
		api.GET("/emulations", emulationH.GetAll)
		api.GET("/emulations/:id", emulationH.GetById)
		api.GET("/users/:id", userH.GetById)

		// Only for users
		auth := api.Group("/", AuthMiddleware("user"))
		{
			auth.GET("/profile", userH.GetProfile)
			auth.POST("/logout", userH.Logout)

			games := auth.Group("/games")
			{
				games.POST("/", gameH.Create)
				games.PUT("/:id", gameH.Update)
				games.DELETE("/:id", gameH.Delete)
			}

			companies := auth.Group("/companies")
			{
				companies.POST("/", companyH.Create)
				companies.PUT("/:id", companyH.Update)
				companies.DELETE("/:id", companyH.Delete)
			}

			emulations := auth.Group("/emulations")
			{
				emulations.POST("/", emulationH.Create)
				emulations.PUT("/:id", emulationH.Update)
				emulations.DELETE("/:id", emulationH.Delete)
			}

			purchases := auth.Group("/purchases")
			{
				purchases.POST("/", purchaseH.Create)
				purchases.GET("/:id", purchaseH.GetById)
				purchases.DELETE("/:id", purchaseH.Delete)
			}

			reviews := auth.Group("/reviews")
			{
				reviews.POST("/", reviewH.Create)
				reviews.GET("/", reviewH.GetAll)
				reviews.GET("/:id", reviewH.GetById)
				reviews.PUT("/:id", reviewH.Update)
				reviews.DELETE("/:id", reviewH.Delete)
			}

			// Only for admins
			admin := auth.Group("/admin", AuthMiddleware("admin"))
			{
				admin.PATCH("/games/:id/verify", gameH.VerifySwitch)
				admin.DELETE("/games/:id/delete", gameH.Delete)
				admin.GET("/users", userH.GetAll)
				admin.DELETE("/users/:id", userH.Delete)
				admin.GET("/purchases", purchaseH.GetAll)
			}
		}
	}
}
