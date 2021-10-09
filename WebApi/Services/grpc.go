package Services

import (
	"WebApi/Pb/user"
	"fmt"
	"google.golang.org/grpc"
)

var UserGrpc user.UserClient

func init() {
	conn, err := grpc.Dial("172.20.3.111:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	UserGrpc = user.NewUserClient(conn)
}
