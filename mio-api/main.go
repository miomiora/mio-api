package main

import (
	"mio-api/routes"
	"mio-api/server"
)

func main() {
	// 启动rpc服务器
	server.StartRPC()
	// 启动gin服务器
	routes.StartGin()
}
