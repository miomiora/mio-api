package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mio-api/config"
	"mio-api/middleware"
)

var r *gin.Engine
var apiGroup *gin.RouterGroup

func init() {
	r = gin.Default()
	//全局中间件
	r.Use(middleware.Cors, middleware.RefreshToken)
	apiGroup = r.Group("api")

	addUserRouter()
	addInterfaceRouter()
}

func StartGin() {
	err := r.Run(config.Conf.Gin.GetAddr())
	if err != nil {
		fmt.Println("[Error] r.Run " + err.Error())
		return
	}
}
