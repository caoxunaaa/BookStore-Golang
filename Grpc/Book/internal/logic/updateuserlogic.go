package logic

import (
	"Book/model"
	"context"
	"database/sql"
	"time"

	"Book/book"
	"Book/internal/svc"

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

func (l *UpdateUserLogic) UpdateUser(in *book.BookBasicInfoReq) (*book.Reply, error) {
	storeTime, err := time.ParseInLocation("2006-01-02 15:04:05", in.StorageTime, time.Local)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.BookBasicInfoModel.Update(model.BookBasicInfo{
		Id:          in.Id,
		Name:        in.Name,
		Author:      in.Author,
		Image:       in.Image,
		StorageTime: sql.NullTime{Time: storeTime},
	})
	if err != nil {
		return nil, err
	}
	return &book.Reply{Ok: true}, nil
}
