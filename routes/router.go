package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    gin.SetMode(gin.DebugMode) 
    router := gin.Default()
	RegisterDeviceRoutes(router)  
    return router
}
