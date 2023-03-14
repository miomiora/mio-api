package routes

import (
	"mio-api/api"
	"mio-api/middleware"
)

func addUserRouter() {
	userApi := api.RootApi.UserApi

	userRouter := apiGroup.Group("user")
	{
		/*
			login		登录
			register	注册
			logout		登出
		*/
		userRouter.POST("login", userApi.UserLogin)
		userRouter.POST("register", userApi.UserRegister)
		userRouter.POST("logout", userApi.UserLogout)

		/*
			普通用户都可访问
			current		根据Token获取的id进行查找用户
			update		普通用户修改自己的信息
			change		普通用户修改自己的密码
		*/
		userRouter.GET("current", middleware.AuthUser, userApi.GetCurrentUser)
		userRouter.PUT("update", middleware.AuthUser, userApi.UpdateUserBySelf)
		userRouter.PUT("change", middleware.AuthUser, userApi.ChangePasswordBySelf)

		/*
			下列需要管理员权限
			list/:num/:page		获取全部的用户(分页)
			search/:id			获取单个用户
			delete/:id			删除指定用户
			update/:id			更新指定用户
			change/:id			修改指定用户的密码
		*/
		userRouter.GET("list/:num/:page", middleware.AuthAdmin, userApi.GetUserList)
		userRouter.GET("search/:id", middleware.AuthAdmin, userApi.GetUserById)
		userRouter.POST("delete/:id", middleware.AuthAdmin, userApi.DeleteUserById)
		userRouter.PUT("update/:id", middleware.AuthAdmin, userApi.UpdateUserById)
		userRouter.PUT("change/:id", middleware.AuthAdmin, userApi.ChangePasswordById)
	}
}
