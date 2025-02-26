package main

import (
	"github.com/leofideliss/devices/routes"
)

func main() {
    
    router:=routes.SetupRouter()

    router.Run()
}
