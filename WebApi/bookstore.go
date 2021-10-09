package main

import (
	"WebApi/Pb/user"
	"WebApi/Router"
	"WebApi/Services"
	"context"
	"fmt"
)

func main() {
	// 调用服务端函数
	res, err := Services.UserGrpc.FindAllUser(context.Background(), &user.Request{})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	fmt.Println(res)

	r := Router.Init()
	if err := r.Run("0.0.0.0:8002"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
