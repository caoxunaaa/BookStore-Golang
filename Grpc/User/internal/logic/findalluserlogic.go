package logic

import (
	"User/internal/svc"
	"User/user"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindAllUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllUserLogic {
	return &FindAllUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindAllUserLogic) FindAllUser(in *user.Request) (*user.UsersInfoReply, error) {
	tmp, err := l.svcCtx.Model.FindAll()
	if err != nil {
		return nil, err
	}
	var res user.UsersInfoReply
	res.UsersInfo = make([]*user.UserInfoReply, 0)
	for i := 0; i < len(tmp); i++ {
		res.UsersInfo = append(res.UsersInfo, &user.UserInfoReply{
			Id:       tmp[i].Id,
			Username: tmp[i].Username,
			//Password: tmp[i].Password,
			Nickname: tmp[i].Nickname,
			Phone:    tmp[i].Phone,
			Email:    tmp[i].Email,
		})
	}
	return &res, nil
}
