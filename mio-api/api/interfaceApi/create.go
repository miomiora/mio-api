package interfaceApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mio-api/model"
	"mio-api/utils"
	"net/http"
	"strconv"
)

func (InterfaceApi) CreateInterface(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "获取用户身份失败！"))
		return
	}
	var getInterface model.NewInterface
	err := c.ShouldBindJSON(&getInterface)
	if err != nil {
		fmt.Println("[Error api interface] CreateInterface c.ShouldBindJSON ", err.Error())
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "新建接口的参数有误！"))
		return
	}

	atoi, err := strconv.Atoi(userId.(string))
	if err != nil {
		fmt.Println("[Error api interface] CreateInterface strconv.Atoi ", err.Error())
		return
	}

	var interfaceInfo = &model.InterfaceInfo{
		Name:           getInterface.Name,
		UserId:         uint(atoi),
		Description:    getInterface.Description,
		Url:            getInterface.Url,
		Method:         getInterface.Method,
		RequestHeader:  getInterface.RequestHeader,
		ResponseHeader: getInterface.ResponseHeader,
		Status:         false,
	}

	affected := db.Save(interfaceInfo).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "创建接口失败！"))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseOK(getInterface))
}
