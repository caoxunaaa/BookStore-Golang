package logic

import (
	"context"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindOneBookContentByBookIdAndChapterNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneBookContentByBookIdAndChapterNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneBookContentByBookIdAndChapterNumLogic {
	return &FindOneBookContentByBookIdAndChapterNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOneBookContentByBookIdAndChapterNumLogic) FindOneBookContentByBookIdAndChapterNum(in *book.BookContentReq) (*book.BookContentReply, error) {
	rep, err := l.svcCtx.BookContentModel.FindOneByBookIdChapterNum(in.BookId, in.ChapterNum)
	if err != nil {
		return nil, err
	}

	return &book.BookContentReply{
		Id:             rep.Id,
		ChapterNum:     rep.ChapterNum,
		ChapterName:    rep.ChapterName,
		ChapterContent: rep.ChapterContent,
		CreateTime:     rep.CreateTime.Time.Format("2006-01-02 15:04:05"),
	}, nil
}
