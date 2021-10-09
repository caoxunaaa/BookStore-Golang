package main

import (
	"context"
	"flag"
	"fmt"

	"User/internal/config"
	"User/internal/server"
	"User/internal/svc"
	"User/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Println(srv.FindAllUser(context.Background(), &user.Request{}))
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
