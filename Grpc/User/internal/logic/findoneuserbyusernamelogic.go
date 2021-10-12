package logic

import (
	"context"

	"User/internal/svc"
	"User/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindOneUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneUserByUsernameLogic {
	return &FindOneUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOneUserByUsernameLogic) FindOneUserByUsername(in *user.UsernameReq) (*user.UserInfoReply, error) {
	res, err := l.svcCtx.Model.FindOneByUsername(in.Username)
	if err != nil {
		return nil, err
	}
	return &user.UserInfoReply{
		Id:       res.Id,
		Username: res.Username,
		//Password: res.Password,
		Nickname: res.Nickname,
		Phone:    res.Phone,
		Email:    res.Email,
	}, nil
}
