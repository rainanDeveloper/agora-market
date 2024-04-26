package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rainanDeveloper/agora-market/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Title = "Agora Market API"
	api := router.Group(basePath)
	{
		api.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Ok",
			})
		})
	}
	router.GET("/docs", func (context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/docs/index.html");
	})
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
