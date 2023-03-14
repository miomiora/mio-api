package routes

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine
var apiGroup *gin.RouterGroup

func init() {
	r := gin.Default()
	//全局中间件
	apiGroup = r.Group("api")
	addNameRouter()
	addSentenceRouter()
	Router = r
}
