package interfaceApi

import (
	"github.com/gin-gonic/gin"
	"mio-api/model"
	"mio-api/utils"
	"net/http"
)

func (InterfaceApi) DeleteInterface(c *gin.Context) {
	// 判断id是否合法
	id := c.Param("id")
	matched := utils.MatchString(`^[0-9]*$`, id)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "id必须为数字！"))
	}

	// 合法就删除
	affected := db.Delete(&model.InterfaceInfo{}, id).RowsAffected
	// 用户已经不存在了
	if affected == 0 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "该接口已经不存在！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil))
}
