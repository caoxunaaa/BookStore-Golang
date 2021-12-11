package logic

import (
	"context"
	"fmt"

	"User/internal/svc"
	"User/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindOneUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneUserByIdLogic {
	return &FindOneUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOneUserByIdLogic) FindOneUserById(in *user.IdReq) (*user.UserInfoReply, error) {
	res, err := l.svcCtx.Model.FindOne(in.Id)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return &user.UserInfoReply{
		Id:       res.Id,
		Username: res.Username,
		//Password: res.Password,
		Nickname: res.Nickname,
		Phone:    res.Phone,
		Email:    res.Email,
	}, nil
}
