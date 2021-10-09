package logic

import (
	"User/model"
	"context"

	"User/internal/svc"
	"User/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.Reply, error) {
	var err error
	_, err = l.svcCtx.Model.FindOne(in.Id)
	if err == model.ErrNotFound {
		return &user.Reply{Ok: false, Code: "id is not exist"}, nil
	}

	err = l.svcCtx.Model.Update(model.UserInfo{
		Id:       in.Id,
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
		Phone:    in.Phone})
	if err != nil {
		return nil, err
	}
	return &user.Reply{Ok: true, Code: "Update success"}, nil
}
