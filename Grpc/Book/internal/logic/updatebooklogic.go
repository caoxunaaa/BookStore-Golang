package logic

import (
	"Book/model"
	"context"
	"time"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBookLogic {
	return &UpdateBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBookLogic) UpdateBook(in *book.BookBasicInfoReq) (*book.Reply, error) {
	storeTime, err := time.Parse("2006-01-02 15:04:05", in.StorageTime)
	//fmt.Println(storeTime)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.BookBasicInfoModel.Update(model.BookBasicInfo{
		Id:            in.Id,
		Name:          in.Name,
		Author:        in.Author,
		Image:         in.Image,
		StorageUserId: in.StorageUserId,
		StorageTime:   storeTime,
	})
	if err != nil {
		return nil, err
	}
	return &book.Reply{Ok: true}, nil
}
