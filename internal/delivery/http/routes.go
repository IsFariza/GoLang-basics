package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, gameH *GameHandler, companyH *CompanyHandler, emulationH *EmulationHandler, userH *UserHandler, purchaseH *PurchaseHandler, reviewH *ReviewHandler) {
	api := r.Group("/api/v1")
	{
		api.POST("/signup", userH.SignUp)
		api.POST("/login", userH.Login)

		api.GET("/games", gameH.GetAll)
		api.GET("/games/:id", gameH.GetById)
		api.GET("/:id/reviews", gameH.GetReviewsByGameId)
		api.GET("/companies", companyH.GetAll)
		api.GET("/companies/:id", companyH.GetById)
		api.GET("/emulations", emulationH.GetAll)
		api.GET("/emulations/:id", emulationH.GetById)

		// Only for users
		auth := api.Group("/", AuthMiddleware("user"))
		{
			auth.POST("/games", gameH.Create)
			auth.PUT("/games/:id", gameH.Update)
			auth.DELETE("/games/:id", gameH.Delete)

			auth.POST("/companies", companyH.Create)
			auth.PUT("/companies/:id", companyH.Update)
			auth.DELETE("/companies/:id", companyH.Delete)

			auth.POST("/emulations", emulationH.Create)
			auth.PUT("/emulations/:id", emulationH.Update)
			auth.DELETE("/emulations/:id", emulationH.Delete)

			auth.POST("/", purchaseH.Create)
			auth.GET("/", purchaseH.GetAll)
			auth.GET("/:id", purchaseH.GetById)
			auth.DELETE("/:id", purchaseH.Delete)

			auth.POST("/", reviewH.Create)
			auth.GET("/", reviewH.GetAll)
			auth.GET("/:id", reviewH.GetById)
			auth.PUT("/:id", reviewH.Update)
			auth.DELETE("/:id", reviewH.Delete)

			// Only for admins
			admin := auth.Group("/admin", AuthMiddleware("admin"))
			{
				admin.PATCH("/games/:id/approve", gameH.Approve)
				admin.GET("/users", userH.GetAll)
				admin.DELETE("/users/:id", userH.Delete)
			}
		}
	}
}
