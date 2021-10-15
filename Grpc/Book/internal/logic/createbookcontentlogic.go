package logic

import (
	"Book/model"
	"context"
	"time"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateBookContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBookContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBookContentLogic {
	return &CreateBookContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBookContentLogic) CreateBookContent(in *book.BookContentReq) (*book.Reply, error) {
	createTime, err := time.ParseInLocation("2006-01-02", in.CreateTime, time.Local)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.BookContentModel.Insert(model.BookContent{
		BookId:         in.BookId,
		ChapterNum:     in.ChapterNum,
		ChapterName:    in.ChapterName,
		ChapterContent: in.ChapterContent,
		CreateTime:     createTime,
	})
	if err != nil {
		return nil, err
	}
	return &book.Reply{Ok: true, Message: in.ChapterName + "上传成功"}, nil
}
