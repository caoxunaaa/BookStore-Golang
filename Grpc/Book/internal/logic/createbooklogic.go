package logic

import (
	"Book/model"
	"context"
	"time"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBookLogic {
	return &CreateBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBookLogic) CreateBook(in *book.BookBasicInfoReq) (*book.Reply, error) {
	storeTime, err :=time.ParseInLocation("2006-01-02", in.StorageTime,time.Local)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.Model.Insert(model.BookBasicInfo{
		Name: in.Name,
		Author: in.Author,
		Image: in.Image,
		StorageTime: storeTime,
	})
	if err != nil {
		return nil, err
	}
	return &book.Reply{Ok: true, Message: in.Name+"入库"}, nil
}
