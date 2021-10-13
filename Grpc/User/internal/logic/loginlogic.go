package logic

import (
	"User/model"
	"context"

	"User/internal/svc"
	"User/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.UserInfoReply, error) {
	var err error
	var rep = new(model.UserInfo)
	if in.Username != "" {
		rep, err = l.svcCtx.Model.FindOneByUsername(in.Username)
	} else if in.Email != "" {
		rep, err = l.svcCtx.Model.FindOneByEmail(in.Email)
	} else if in.Phone != "" {
		rep, err = l.svcCtx.Model.FindOneByPhone(in.Phone)
	} else {
		return &user.UserInfoReply{}, nil
	}
	if err != nil {
		return nil, err
	}
	if rep.Password == in.Password {
		return &user.UserInfoReply{
			Id:       rep.Id,
			Username: rep.Username,
			Nickname: rep.Nickname,
			//Password: rep.Password,
			Email: rep.Email,
			Phone: rep.Phone,
		}, nil
	}
	return &user.UserInfoReply{}, nil
}
