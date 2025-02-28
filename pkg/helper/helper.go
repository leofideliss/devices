package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
    Message string `json:"message"`
    Success bool `json:"success"`
    Data any `json:"data,omitempty"`
}

func HandleResponseJson( c *gin.Context , data any , message string , status int , success bool){
    c.JSON(status,Response{ Message: message , Success : success , Data: data})
}

