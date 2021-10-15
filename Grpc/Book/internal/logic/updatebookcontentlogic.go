package logic

import (
	"Book/model"
	"context"
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
	createTime, err := time.ParseInLocation("2006-01-02", in.CreateTime, time.Local)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.BookContentModel.Update(model.BookContent{
		BookId:         in.BookId,
		ChapterNum:     in.ChapterNum,
		ChapterName:    in.ChapterName,
		ChapterContent: in.ChapterContent,
		CreateTime:     createTime,
	})
	if err != nil {
		return nil, err
	}
	return &book.Reply{Ok: true}, nil
}
