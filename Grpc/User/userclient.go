package main

import (
	"User/user"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("172.20.3.111:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务端失败: ", err)
		return
	}
	defer conn.Close()

	// 新建一个客户端
	c := user.NewUserClient(conn)

	// 调用服务端函数
	r, err := c.FindAllUser(context.Background(), &user.Request{})
	if err != nil {
		fmt.Println("调用服务端代码失败: ", err)
		return
	}

	fmt.Println("调用成功: ", r)
}
