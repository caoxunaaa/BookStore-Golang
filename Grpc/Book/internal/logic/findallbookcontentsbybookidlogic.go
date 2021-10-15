package logic

import (
	"Book/model"
	"context"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindAllBookContentsByBookIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllBookContentsByBookIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllBookContentsByBookIdLogic {
	return &FindAllBookContentsByBookIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//   book_content
func (l *FindAllBookContentsByBookIdLogic) FindAllBookContentsByBookId(in *book.BookContentReq) (*book.BookContentsReply, error) {
	reps, err := l.svcCtx.BookContentModel.FindAllBookContentsByBookId(in.BookId)
	if err != nil {
		return nil, err
	}
	f := func(t []*model.BookContent) []*book.BookContentReply {
		var res = make([]*book.BookContentReply, 0)
		for i := 0; i < len(t); i++ {
			res = append(res, &book.BookContentReply{
				Id:             t[i].Id,
				BookId:         t[i].BookId,
				ChapterNum:     t[i].ChapterNum,
				ChapterName:    t[i].ChapterName,
				ChapterContent: t[i].ChapterContent,
				CreateTime:     t[i].CreateTime.Format("2006-01-02"),
			})
		}
		return res
	}
	return &book.BookContentsReply{BookContentsReply: f(reps)}, nil
}
