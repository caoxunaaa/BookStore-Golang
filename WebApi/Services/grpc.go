package Services

import (
	"WebApi/Pb/user"
	"google.golang.org/grpc"
)

var UserGrpc user.UserClient

func GrpcInit() error {
	conn, err := grpc.Dial(C.UserRpc.Host, grpc.WithInsecure())
	if err != nil {
		return err
	}
	UserGrpc = user.NewUserClient(conn)
	return nil
}
