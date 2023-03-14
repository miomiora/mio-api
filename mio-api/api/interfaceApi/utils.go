package interfaceApi

import (
	"github.com/gin-gonic/gin"
	"mio-api/model"
	"mio-api/utils"
	"net/http"
	"strconv"
)

func setStatus(c *gin.Context, status bool) {
	id := utils.IsNumber(c.Param("id"))
	if id == -1 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "id错误！"))
		return
	}
	affected := db.Take(&model.InterfaceInfo{}, id).Select("status").Updates(model.InterfaceInfo{
		Status: status,
	}).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "更新失败！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil, "发布成功！"))
}

func getInterfaceList(c *gin.Context, flag bool) {
	// 获取num和page的参数，并验证
	numParam := c.Param("num")
	matched := utils.MatchString(`^[0-9]*$`, numParam)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "num必须为数字！"))
		return
	}
	pageParam := c.Param("page")
	matched = utils.MatchString(`^[0-9]*$`, pageParam)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "page必须为数字！"))
		return
	}
	// 将num和page转化为整数
	num, err := strconv.Atoi(numParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "转化为数字失败！"))
		return
	}
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "转化为数字失败！"))
		return
	}
	var interfaceList []model.QueryInterface
	offset := (page - 1) * num

	var affected int64
	var total int64
	if flag {
		affected = db.Model(&model.InterfaceInfo{}).Count(&total).Limit(num).Offset(offset).Scan(&interfaceList).RowsAffected
	} else {
		affected = db.Model(&model.InterfaceInfo{}).Where("status = ?", true).Count(&total).Limit(num).Offset(offset).Scan(&interfaceList).RowsAffected
	}

	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "数据库中已无接口！"))
		return
	}

	res := gin.H{
		"table": interfaceList,
		"total": total,
	}
	c.JSON(http.StatusOK, utils.ResponseOK(res))
}
