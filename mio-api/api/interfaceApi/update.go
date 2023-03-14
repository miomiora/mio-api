package interfaceApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mio-api/model"
	"mio-api/utils"
	"net/http"
	"strconv"
)

func (InterfaceApi) UpdateInterface(c *gin.Context) {
	//验证id是否合法
	id := c.Param("id")
	matched := utils.MatchString(`^[0-9]*$`, id)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "id必须为数字！"))
	}

	// 获取前端发送的接口信息
	var updateInterface model.UpdateInterface
	err := c.ShouldBindJSON(&updateInterface)
	if err != nil {
		fmt.Println("[Error Interface update] UpdateInterface c.ShouldBindJSON ", err.Error())
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "修改的参数出错！"))
		return
	}

	// 判断请求参数的id是否和获取到的接口id一致，不一致则直接返回
	if id != strconv.Itoa(int(updateInterface.ID)) {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "需要修改的接口参数和请求不合法！"))
		return
	}

	// 更新接口的信息
	affected := db.Take(&model.InterfaceInfo{}, updateInterface.ID).Updates(model.InterfaceInfo{
		Name:           updateInterface.Name,
		UserId:         updateInterface.UserId,
		Description:    updateInterface.Description,
		Url:            updateInterface.Url,
		Method:         updateInterface.Method,
		ResponseHeader: updateInterface.ResponseHeader,
		RequestHeader:  updateInterface.RequestHeader,
		Status:         updateInterface.Status,
		RequestParams:  updateInterface.RequestParams,
	}).RowsAffected

	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "更新失败！"))
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil, "修改接口信息成功！"))
}

func (InterfaceApi) OnlineInterface(c *gin.Context) {
	// 测试接口是否可用
	//client := mioApiClient.NewClient("miomiora", "pass", )
	//request := client.DoRequest()
	//if request == nil {
	//	c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "接口不可用"))
	//	return
	//}
	setStatus(c, true)
}

func (InterfaceApi) OfflineInterface(c *gin.Context) {
	setStatus(c, false)
}
