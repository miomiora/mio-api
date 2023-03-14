package userApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mio-api/model"
	"mio-api/utils"
	"net/http"
	"strconv"
)

// GetUserById
//
//	@Description: 获取指定id的用户, 需要接受请求参数:id
//	@param c
func (UserApi) GetUserById(c *gin.Context) {
	// 验证id是否合法
	id := utils.IsNumber(c.Param("id"))
	if id == -1 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "id错误！"))
		return
	}

	var userDTO model.UserInfo
	affected := db.Take(&model.User{}, id).Scan(&userDTO).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "该用户不存在！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(userDTO))
}

// GetUserList
//
//	@Description: 获取全部用户，需要接收请求参数: num一页的数量，page当前的页数
//	@param c
func (UserApi) GetUserList(c *gin.Context) {
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

	var userList []model.UserInfo
	offset := (page - 1) * num
	affected := db.Limit(num).Offset(offset).Model(&model.User{}).Scan(&userList).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "数据库中没有用户！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(userList, fmt.Sprintf("查找到了%v个用户！", affected)))
}

// GetCurrentUser
//
//	@Description: 获取当前用户,无需前端传参数,将从token中取得当前的用户
//	@param c
func (UserApi) GetCurrentUser(c *gin.Context) {
	userId, exists := c.Get("userId")
	var user model.UserInfo
	if !exists {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "获取用户身份失败！"))
		return
	}
	affected := db.Take(&model.User{}, userId).Scan(&user).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "该用户不存在！"))
		return
	}
	// 能走到这个函数，说明已经验证过用户存在，直接给前端返回200
	c.JSON(http.StatusOK, utils.ResponseOK(user))
}
