package routes

import (
	"mio-api-interface/api"
	"mio-api-interface/middleware"
)

func addNameRouter() {
	nameRouter := apiGroup.Group("user")
	{
		nameRouter.GET("name", middleware.ValidHeader, api.RootApi.GetUserName)
		nameRouter.GET("get", api.RootApi.GetName)
		nameRouter.POST("post", api.RootApi.PostName)
	}
}
