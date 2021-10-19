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

type UpdateBookContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBookContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBookContentLogic {
	return &UpdateBookContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBookContentLogic) UpdateBookContent(in *book.BookContentReq) (*book.Reply, error) {
	createTime, err := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.BookContentModel.Update(model.BookContent{
		BookId:         in.BookId,
		ChapterNum:     in.ChapterNum,
		ChapterName:    in.ChapterName,
		ChapterContent: in.ChapterContent,
		CreateTime:     sql.NullTime{Time: createTime, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &book.Reply{Ok: true}, nil
}
