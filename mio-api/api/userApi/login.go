package userApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mio-api/model"
	"mio-api/utils"
	"net/http"
)

// UserLogin
//
//	@Description: 用户登录，需要参数：user_account、user_password
//	@param c
func (UserApi) UserLogin(c *gin.Context) {
	var userDTO model.UserInfo
	var user model.User
	// 获取用户登录信息, 同时校验是否为空, 以及长度是否合法
	result := bindContextJson(c, &user)
	if !result {
		return
	}

	// 帐号是否合法(字母开头，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]*$
	matched := utils.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, user.UserAccount)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "账号不合法！"))
		return
	}
	// 加密密码
	password := encryptString(user.UserPassword)

	// 查询数据库中是否存在该用户，并且同时把取出来的数据存入userDTO中
	affected := db.Take(&model.User{},
		"user_account = ? and user_password = ?", user.UserAccount, password).
		Scan(&userDTO).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "账号不存在！"))
		return
	}

	// 记录用户的登录状态, 使用redis+token
	token := uuid.NewString()
	tokenKey := utils.TokenPrefix + token
	// 存入redis, 并且把用户ip存入redis
	err := client.HSet(ctx, tokenKey, "id", userDTO.ID, "client_ip", c.ClientIP()).Err()

	if err != nil {
		fmt.Println("[api UserLogin err] Conn.Do HSET : " + err.Error())
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.RedisError, "存储Token失败！"))
		return
	}
	// 设置有效期
	err = client.Expire(ctx, tokenKey, utils.TokenTimeout).Err()
	if err != nil {
		fmt.Println("[api UserLogin err] Conn.Do EXPIRE : " + err.Error())
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.RedisError, "设置Token有效期失败！"))
		return
	}

	// 封装user和token
	res := &gin.H{
		"user":  userDTO,
		"token": token,
	}
	c.JSON(http.StatusOK, utils.ResponseOK(res))
}

// UserLogout
//
//	@Description: 用户登出,无需前端传参数,将从token中取得当前的用户
//	@param c
func (UserApi) UserLogout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	tokenKey := utils.TokenPrefix + token
	if token != "" {
		err := client.HDel(ctx, tokenKey, "client_ip", "id").Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.RedisError, "用户Token已失效！"))
		}
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil))
}
