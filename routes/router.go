package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    
	RegisterDeviceRoutes(router)  

    return router
}
