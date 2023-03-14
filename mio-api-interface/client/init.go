package client

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mio-api-interface/config"
	pb "mio-api-interface/proto"
)

var RPCClient pb.MioApiClient

func init() {
	dial, err := grpc.Dial(config.Conf.GRPC.GetAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("客户端链接失败！")
	}

	// 建立链接
	RPCClient = pb.NewMioApiClient(dial)
}
