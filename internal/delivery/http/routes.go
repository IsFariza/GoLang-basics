package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(companyH *CompanyHandler) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		companies := api.Group("/companies")
		{
			companies.POST("/", companyH.Create)
			companies.GET("/", companyH.GetAll)
			companies.GET("/:id", companyH.GetById)
			companies.PUT("/:id", companyH.Update)
			companies.DELETE("/:id", companyH.Delete)
		}
	}
	return r
}
