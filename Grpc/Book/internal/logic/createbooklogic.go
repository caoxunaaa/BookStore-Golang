package logic

import (
	"Book/model"
	"context"
	"database/sql"
	"fmt"
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
	storeTime, err := time.Parse("2006-01-02 15:04:05", in.StorageTime)
	fmt.Println(storeTime)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.BookBasicInfoModel.Insert(model.BookBasicInfo{
		Name:          in.Name,
		Author:        in.Author,
		Image:         in.Image,
		StorageUserId: in.StorageUserId,
		StorageTime:   sql.NullTime{Time: storeTime, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &book.Reply{Ok: true, Message: in.Name + "入库"}, nil
}
