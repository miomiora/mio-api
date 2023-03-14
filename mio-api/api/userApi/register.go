package userApi

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mio-api/model"
	"mio-api/utils"
	"net/http"
)

// UserRegister
//
//	@Description: 用户注册, 需要接受json内容:user_account、user_password、check_password
//	@param c
func (UserApi) UserRegister(c *gin.Context) {
	var userRegister model.UserRegister
	// 获取用户注册信息, 同时校验是否为空, 以及长度是否合法
	result := bindContextJson(c, &userRegister)
	if !result {
		return
	}
	// 帐号是否合法(字母开头，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]*$
	matched := utils.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, userRegister.UserAccount)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "账号不合法！"))
		return
	}
	// 密码(以字母开头，只能包含字母、数字和下划线)：^[a-zA-Z]\w*$    \w = [a-zA-Z0-9_]

	// 账户是否重复
	exist := isUserAccountExist(c, userRegister.UserAccount)
	if exist {
		return
	}

	// 密码加密
	password := encryptString(userRegister.UserPassword)
	// accessKey生成
	accessKey := encryptString(userRegister.UserAccount + uuid.NewString())
	secretKey := encryptString(password + uuid.NewString())

	var userDTO model.UserInfo

	// 插入数据
	user := &model.User{
		UserPassword: password,
		UserAccount:  userRegister.UserAccount,
		AccessKey:    accessKey,
		SecretKey:    secretKey,
	}
	affected := db.Save(user).Scan(&userDTO).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "注册用户失败！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(userDTO))
}
