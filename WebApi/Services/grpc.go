package Services

import (
	"WebApi/Pb/book"
	"WebApi/Pb/user"
	"google.golang.org/grpc"
)

var Grpc *GrpcContext

type GrpcContext struct {
	UserGrpc user.UserClient
	BookGrpc book.BookClient
}

func GrpcInit() *GrpcContext {
	var g GrpcContext
	conn, err := grpc.Dial(C.UserRpc.Host, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	g.UserGrpc = user.NewUserClient(conn)

	conn, err = grpc.Dial(C.BookRpc.Host, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	g.BookGrpc = book.NewBookClient(conn)

	return &g
}
