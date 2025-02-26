package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leofideliss/devices/handler"
	"github.com/leofideliss/devices/repository"
	"github.com/leofideliss/devices/service"
)

// Função para registrar rotas de dispositivos
func RegisterDeviceRoutes(router *gin.Engine) {
    deviceService := service.NewDeviceService(repository.DB) 
    deviceHandler := handler.NewDeviceHandler(deviceService)
    
    router.GET("/:id", deviceHandler.GetDevice)

}
