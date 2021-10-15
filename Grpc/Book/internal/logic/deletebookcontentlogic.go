package logic

import (
	"context"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteBookContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBookContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBookContentLogic {
	return &DeleteBookContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBookContentLogic) DeleteBookContent(in *book.BookContentReq) (*book.Reply, error) {
	if err := l.svcCtx.BookContentModel.Delete(in.Id); err != nil {
		return nil, err
	}
	return &book.Reply{Ok: true}, nil
}
