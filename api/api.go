package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rainanDeveloper/agora-market/api/router"
)

func Initialize() *gin.Engine {
	api := gin.Default()
	router.InitializeRoutes(api)
	return api
}
