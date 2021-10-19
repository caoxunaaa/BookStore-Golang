package logic

import (
	"Book/model"
	"context"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindBooksByStorageUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindBooksByStorageUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindBooksByStorageUserIdLogic {
	return &FindBooksByStorageUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindBooksByStorageUserIdLogic) FindBooksByStorageUserId(in *book.BookBasicInfoReq) (*book.BooksBasicInfoReply, error) {
	reps, err := l.svcCtx.BookBasicInfoModel.FindBooksByStorageUserId(in.StorageUserId)
	if err != nil {
		return nil, err
	}
	f := func(t []*model.BookBasicInfo) []*book.BookBasicInfoReply {
		var res = make([]*book.BookBasicInfoReply, 0)
		for i := 0; i < len(t); i++ {
			res = append(res, &book.BookBasicInfoReply{
				Id:            t[i].Id,
				Name:          t[i].Name,
				Author:        t[i].Author,
				Image:         t[i].Image,
				StorageUserId: t[i].StorageUserId,
				StorageTime:   t[i].StorageTime.Time.Format("2006-01-02 15:04:05"),
			})
		}
		return res
	}
	return &book.BooksBasicInfoReply{BooksBasicInfo: f(reps)}, nil
}
