package routes

import (
	"mio-api/api"
	"mio-api/middleware"
)

func addInterfaceRouter() {
	interfaceApi := api.RootApi.InterfaceApi

	interfaceRouter := apiGroup.Group("interface")
	{
		interfaceRouter.POST("/invoke/:id", middleware.AuthUser, interfaceApi.InvokeInterface)

		interfaceRouter.GET("/:num/:page", middleware.AuthAdmin, interfaceApi.GetInterfaceList)
		interfaceRouter.GET("/user/:num/:page", middleware.AuthUser, interfaceApi.GetInterfaceListByUser)

		interfaceRouter.POST("/create", middleware.AuthAdmin, interfaceApi.CreateInterface)
		interfaceRouter.POST("/update/:id", middleware.AuthAdmin, interfaceApi.UpdateInterface)
		interfaceRouter.POST("/delete/:id", middleware.AuthAdmin, interfaceApi.DeleteInterface)
		interfaceRouter.POST("/online/:id", middleware.AuthAdmin, interfaceApi.OnlineInterface)
		interfaceRouter.POST("/offline/:id", middleware.AuthAdmin, interfaceApi.OfflineInterface)
	}
}
