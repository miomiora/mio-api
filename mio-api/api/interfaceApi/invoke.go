package interfaceApi

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	mioApiClient "mio-api-client"
	"mio-api/model"
	"mio-api/utils"
	"net/http"
)

func (InterfaceApi) InvokeInterface(c *gin.Context) {

	// 判断id是否合法
	id := c.Param("id")
	matched := utils.MatchString(`^[0-9]*$`, id)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "id必须为数字！"))
	}

	// 拿到请求的参数原始数据
	body, err := c.GetRawData()
	if err != nil {
		fmt.Println("[api interfaceApi err] InvokeInterface io.ReadAll(c.Request.Body) ", err.Error())
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "调用的请求有误！"))
		return
	}

	var invokeInterface model.InvokeInterface
	err = json.Unmarshal(body, &invokeInterface)
	if err != nil {
		fmt.Println("[api interfaceApi err] InvokeInterface c.ShouldBindJSON ", err.Error())
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "调用的请求有误！"))
		return
	}

	// 通过id查找interface的信息
	var interfaceInfo model.InterfaceInfo
	affected := db.Model(&model.InterfaceInfo{}).Where("id = ? and status = ?", id, true).Scan(&interfaceInfo).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "接口不可用！"))
		return
	}
	info, err := json.Marshal(&interfaceInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError))
		return
	}
	log.Println("interfaceInfo: ", string(info))
	// 拿到调用接口的真实客户端，并把body和interfaceInfo传过去
	client, err := mioApiClient.NewClient(body, info)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "调用接口失败！"))
		return
	}
	response := client.DoRequest()
	var res utils.Response
	err = json.Unmarshal(response, &res)
	if err != nil {
		fmt.Println("[Unmarshal context err] ", err.Error())
		return
	}
	if res.Code != 0 {
		c.JSON(http.StatusForbidden, res)
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(res.Data))
}
