package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leofideliss/devices/service"
)

type DeviceHandler struct{
    service *service.DeviceService
}

func NewDeviceHandler(service *service.DeviceService) *DeviceHandler {
    return &DeviceHandler{service: service}
}

func (d *DeviceHandler) GetDevice (c *gin.Context){

}


