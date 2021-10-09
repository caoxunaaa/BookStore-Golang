package logic

import (
	"User/internal/svc"
	"User/model"
	"User/user"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.Reply, error) {
	var err error
	_, err = l.svcCtx.Model.FindOneByUsername(in.Username)
	if err != model.ErrNotFound {
		return &user.Reply{Ok: false, Code: "username have been exist"}, nil
	}
	_, err = l.svcCtx.Model.FindOneByEmail(in.Email)
	if err != model.ErrNotFound {
		return &user.Reply{Ok: false, Code: "email have been exist"}, nil
	}
	_, err = l.svcCtx.Model.FindOneByPhone(in.Phone)
	if err != model.ErrNotFound {
		return &user.Reply{Ok: false, Code: "phone have been exist"}, nil
	}

	if in.Password != in.RepeatPassword {
		return &user.Reply{Ok: false, Code: "the two passwords don't match"}, nil
	}

	_, err = l.svcCtx.Model.Insert(model.UserInfo{
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
		Phone:    in.Phone})
	if err != nil {
		return nil, err
	}

	return &user.Reply{Ok: true, Code: "register success"}, nil
}
