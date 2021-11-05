package Services

import (
	"WebApi/Pb/action"
	"WebApi/Pb/book"
	"WebApi/Pb/user"
	"google.golang.org/grpc"
)

//var Grpc *GrpcContext

type GrpcContext struct {
	UserGrpc   user.UserClient
	BookGrpc   book.BookClient
	ActionGrpc action.ActionClient
}

func GrpcInit(c *Config) *GrpcContext {
	var g GrpcContext
	conn, err := grpc.Dial(c.UserRpc.Host, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	g.UserGrpc = user.NewUserClient(conn)

	conn, err = grpc.Dial(c.BookRpc.Host, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	g.BookGrpc = book.NewBookClient(conn)

	conn, err = grpc.Dial(c.ActionRpc.Host, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	g.ActionGrpc = action.NewActionClient(conn)

	return &g
}
