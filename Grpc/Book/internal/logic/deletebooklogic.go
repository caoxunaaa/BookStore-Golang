package logic

import (
	"context"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBookLogic {
	return &DeleteBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBookLogic) DeleteBook(in *book.BookBasicInfoReq) (*book.Reply, error) {
	if err := l.svcCtx.Model.Delete(in.Id); err!=nil{
		return nil, err
	}
	return &book.Reply{Ok: true}, nil
}
