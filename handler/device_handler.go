package handler

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leofideliss/devices/domain"
	"github.com/leofideliss/devices/pkg/helper"
	"github.com/leofideliss/devices/service"
    "github.com/go-playground/validator/v10"
)

type DeviceHandler struct{
    service *service.DeviceService
}

func NewDeviceHandler(service *service.DeviceService) *DeviceHandler {
    return &DeviceHandler{service: service}
}

func (d *DeviceHandler) GetDevice (c *gin.Context){
    id:=c.Param("id")
    owner:=c.Query("owner")
    hash := fmt.Sprintf("%x",sha256.Sum256([]byte(id + owner)))
    device , err := d.service.GetDevice(context.Background(),hash)
    if err != nil{
        helper.HandleResponseJson(c,map[string]any{"errCode" : "ALR2"},"Fail",http.StatusForbidden,false)
        return
    }
    helper.HandleResponseJson(c,device,"Device consulted",http.StatusOK,true)
    return
}

func (d *DeviceHandler) RegisterDevice (c *gin.Context){
    var request domain.RequestDevice

    if err := c.BindJSON(&request); err != nil{
        helper.HandleResponseJson(c,nil,err.Error(),http.StatusInternalServerError,false)
        return
    }

    validate := validator.New()
    err := validate.Struct(request)
    if err != nil {
        errorMap := make([]string,0)
        for _, e := range err.(validator.ValidationErrors) {
            errorMap = append(errorMap , fmt.Sprintf("Field '%s': %s", e.Field(), e.Tag()))
        }
        helper.HandleResponseJson(c,errorMap,"Validation Error",http.StatusUnprocessableEntity,false)
        return
    }
    
    hash := fmt.Sprintf("%x",sha256.Sum256([]byte(request.DeviceId + request.Owner)))
    
    var device =  domain.Device{
        Id: hash,
        Owner: request.Owner,
        Title: request.Title,
        Metadata: request.Metadata,
        Expires_at: request.Expires_at,
    }

    result,err := d.service.RegisterDevice(context.Background(),&device)
    if err != nil {
        helper.HandleResponseJson(c,map[string]any{"errCode":"ALR1"},"Device already in base",http.StatusForbidden,false)
        return  
    }  
    helper.HandleResponseJson(c,map[string]any{"id" : result.InsertedID},"Device registred",http.StatusCreated,true)
    return
}

func (d *DeviceHandler) DeleteDevice (c *gin.Context) {
    id:=c.Param("id")
    owner:=c.Query("owner")
    hash := fmt.Sprintf("%x",sha256.Sum256([]byte(id + owner)))
    result , err := d.service.DeleteDevice(context.Background(),hash)
    if err == nil {
        if result.DeletedCount != 0 {
            helper.HandleResponseJson(c,nil,"Device deleted",http.StatusOK,true)
            return
        }       
    }
    helper.HandleResponseJson(c,nil,"Error",http.StatusForbidden,false)
    return  
}

func (d *DeviceHandler) UpdateDevice(c *gin.Context){
    id:=c.Param("id")
    owner:=c.Query("owner")
    hash := fmt.Sprintf("%x",sha256.Sum256([]byte(id + owner)))

    var request domain.RequestDevice

    if err := c.BindJSON(&request); err != nil{
        helper.HandleResponseJson(c,nil,"Error Json",http.StatusInternalServerError,false)
        return
    }

    validate := validator.New()
    err := validate.Struct(request)
    if err != nil {
        errorMap := make([]string,0)
        for _, e := range err.(validator.ValidationErrors) {
            errorMap = append(errorMap , fmt.Sprintf("Field '%s': %s", e.Field(), e.Tag()))
        }
        helper.HandleResponseJson(c,errorMap,"Validation Error",http.StatusUnprocessableEntity,false)
        return
    }
    
    var device =  domain.Device{
        Title: request.Title,
        Metadata: request.Metadata,
        Expires_at: request.Expires_at,
        Owner: owner,
    }

    result , err := d.service.UpdateDevice(context.Background(),&device,hash)
    if err != nil {
        helper.HandleResponseJson(c,map[string]any{"errCode":"ALR2"},"Fail",http.StatusForbidden,false)
        return
    }
    helper.HandleResponseJson(c,result.UpsertedID,"Device updated",http.StatusOK,true)
    return
}

func (d *DeviceHandler) ListDevice (c *gin.Context) {
    owner := c.Query("owner")
    limit , _ := strconv.Atoi(c.Query("limit"))
    page , _ := strconv.Atoi(c.Query("page"))
    
    result , err := d.service.ListDevice( context.Background(), owner , limit , page)
    if err != nil {
        helper.HandleResponseJson(c,map[string]any{"errCode":"ALR2"},"Fail",http.StatusForbidden,false)
        return
    }
    helper.HandleResponseJson(c,result,"Devices",http.StatusOK,true)
    return
}
