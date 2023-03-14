package server

import (
	"context"
	"fmt"
	"mio-api/database"
	"mio-api/model"
	pb "mio-api/proto"
)

func (receiver *RPCServer) GetSecretKey(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	var secretKey string
	fmt.Println("get accessKey: ", req.RequestData)
	affected := database.DB.Model(&model.User{}).Where("access_key = ?", req.RequestData).
		Select("secret_key").Scan(&secretKey).RowsAffected
	if affected == 0 {
		return &pb.Response{ResponseData: ""}, nil
	}
	fmt.Println("secretKey: ", secretKey)

	return &pb.Response{ResponseData: secretKey}, nil
}
