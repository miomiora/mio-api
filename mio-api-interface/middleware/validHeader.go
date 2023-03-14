package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mio-api-interface/utils"
	"net/http"
)

func ValidHeader(c *gin.Context) {
	header := c.GetHeader("mio")
	fmt.Println("header: ", header)
	if header == "accept" {
		c.Next()
	}
	c.JSON(http.StatusForbidden, utils.ResponseError(utils.NoAuth, "非法的访问"))
	c.Abort()
}
