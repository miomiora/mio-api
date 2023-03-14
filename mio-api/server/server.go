package server

import (
	"fmt"
	"google.golang.org/grpc"
	"mio-api/config"
	pb "mio-api/proto"
	"net"
)

type RPCServer struct {
	pb.UnimplementedMioApiServer
}

func StartRPC() {
	// 开启grpc端口监听
	// 开启端口
	listen, _ := net.Listen("tcp", config.Conf.GRPC.GetAddr())
	// 创建grpc服务
	newServer := grpc.NewServer()
	pb.RegisterMioApiServer(newServer, &RPCServer{})

	// 启动服务
	go func() {
		err := newServer.Serve(listen)
		if err != nil {
			fmt.Println("[Error] 启动grpc服务器失败！ ", err.Error())
		}
		fmt.Println("gRPC is running in ", config.Conf.GRPC.GetAddr())
	}()
}
