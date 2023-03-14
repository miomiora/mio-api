package main

import (
	"mio-api-interface/config"
	"mio-api-interface/routes"
)

func main() {
	err := routes.Router.Run(config.Conf.Gin.GetAddr())
	if err != nil {
		return
	}
}
