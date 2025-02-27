package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leofideliss/devices/domain"
	"github.com/leofideliss/devices/pkg/helper"
	"github.com/leofideliss/devices/service"
)

type DeviceHandler struct{
    service *service.DeviceService
}

func NewDeviceHandler(service *service.DeviceService) *DeviceHandler {
    return &DeviceHandler{service: service}
}

func (d *DeviceHandler) GetDevice (c *gin.Context){
    device , err := d.service.GetDevice(context.Background(),c.Param("id"))
    helper.HandleResponseJson(c,device,"Device consultado com sucesso!",err,"Não existe registro para esse id")
    return
}

func (d *DeviceHandler) RegisterDevice (c *gin.Context){
    var request domain.RequestDevice

    
    if err := c.BindJSON(&request); err != nil{
        c.JSON(http.StatusInternalServerError,helper.Response{Message:"Erro body to json",Success:false})
        return
    }

    var device =  domain.Device{
        Id: request.DeviceId,
        Owner: request.Owner,
        Title: request.Title,
        Metadata: request.Metadata,
        Expires_at: request.Expires_at,
    }
    
    result,err := d.service.RegisterDevice(context.Background(),&device)
    helper.HandleResponseJson(c,result.InsertedID,"Device cadastrado com sucesso!",err,"Ao cadastrar dispositivo")
    return
}
