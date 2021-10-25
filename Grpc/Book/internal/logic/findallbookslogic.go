package logic

import (
	"Book/book"
	"Book/internal/svc"
	"Book/model"
	"context"
	"github.com/tal-tech/go-zero/core/logx"
)

type FindAllBooksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllBooksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllBooksLogic {
	return &FindAllBooksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindAllBooksLogic) FindAllBooks(in *book.Request) (*book.BooksBasicInfoReply, error) {
	reps, err := l.svcCtx.BookBasicInfoModel.FindAll()
	if err != nil {
		return nil, err
	}
	//fmt.Println(reps)
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
