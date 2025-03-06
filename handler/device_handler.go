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
    _ "github.com/leofideliss/devices/docs"
)

type DeviceHandler struct{
    service *service.DeviceService
}

func NewDeviceHandler(service *service.DeviceService) *DeviceHandler {
    return &DeviceHandler{service: service}
}

// GetDevice godoc
// @Summary Consulta um dispositivo pelo Id
// @Description Retorna as informações do dispositivo
// @Tags Gerenciar Dispositivos
// @Accept json
// @Produce json
// @Param id path string true "ID do dispositivo"
// @Param owner query string false "Owner do dispositivo" 
// @Success 200 {object} helper.Response "Responde com as informações do dispositivo"
// @Router /{id} [get]
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

// RegisterDevice godoc
// @Summary Registra um dispositivo
// @Description Retorna o _id do dispostivo cadastrado
// @Tags Gerenciar Dispositivos
// @Accept json
// @Produce json
// @Param device body domain.RequestDeviceSwagger true "Dados do dispositivo"
// @Success 200 {object} helper.Response "Responde com o id do dispositivo cadastrado"
// @Router /register [post]
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

// DeleteDevice godoc
// @Summary Deleta um dispositivo
// @Description Retorna se o dispositivo foi deletado
// @Tags Gerenciar Dispositivos
// @Accept json
// @Produce json
// @Param id path string true "ID do dispositivo"
// @Param owner query string false "Owner do dispositivo" 
// @Success 200 {object} helper.Response "Responde se o dispositivo foi deletado com sucesso"
// @Router /{id} [delete]
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

// UpdateDevice godoc
// @Summary Atualiza um dispositivo
// @Description Retorna o _id do dispostivo atualizado
// @Tags Gerenciar Dispositivos
// @Accept json
// @Produce json
// @Param id path string true "ID do dispositivo"
// @Param owner query string false "Owner do dispositivo" 
// @Param device body domain.RequestDeviceSwagger true "do dispositivo"
// @Success 200 {object} helper.Response "Responde com o id do dispositivo cadastrado"
// @Router /{id} [patch]
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

// ListDevice godoc
// @Summary Lista todos os dispositivos 
// @Description Retorna uma lista com os dispositivos
// @Tags Gerenciar Dispositivos
// @Accept json
// @Produce json
// @Param owner query string false "Owner do dispositivo" 
// @Param limit query string false "Limite de registros consultados" 
// @Param page query string false "Pagina atual consultada" 
// @Success 200 {object} helper.Response "Responde com o id do dispositivo cadastrado"
// @Router /list [get]
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
