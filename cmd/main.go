package main

import (
	"github.com/leofideliss/devices/routes"
    _ "github.com/leofideliss/devices/docs"

    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"

)

// @title Microsserviço de Dispositivos
// @version 1.0
// @description API para o gerenciamento de dispositivos

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8888

// @contact.name   Leonardo Fidelis
// @contact.url    https://github.com/leofideliss

func main() {
    router:=routes.SetupRouter()
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.Run(":8888")
}
