package api

import (
	"mio-api/api/interfaceApi"
	"mio-api/api/userApi"
)

type Api struct {
	UserApi      userApi.UserApi
	InterfaceApi interfaceApi.InterfaceApi
}

var RootApi = new(Api)
