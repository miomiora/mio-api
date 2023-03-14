package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mio-api-interface/utils"
	"net/http"
)

type GetUserNameApi struct {
	UserName string `json:"user_name"`
}

func (Api) GetUserName(c *gin.Context) {
	body := c.GetHeader("body")
	var userInfo GetUserNameApi
	fmt.Println("body: ", body)
	err := json.Unmarshal([]byte(body), &userInfo)
	if err != nil {
		fmt.Println("[api userApi err] json.Unmarshal ", err.Error())
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "参数有误！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(userInfo.UserName))
}

func (Api) GetName(c *gin.Context) {
	c.JSON(200, "get miomiora")
}

func (Api) PostName(c *gin.Context) {
	c.JSON(200, "post miomiora")
}
