package logic

import (
	"context"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindBookInventoryByBookIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindBookInventoryByBookIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindBookInventoryByBookIdLogic {
	return &FindBookInventoryByBookIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindBookInventoryByBookIdLogic) FindBookInventoryByBookId(in *book.BookInventoryReq) (*book.BookInventoryResp, error) {
	resp, err := l.svcCtx.BookInventoryModel.FindOneByBookId(in.BookId)
	if err != nil {
		return nil, err
	}
	return &book.BookInventoryResp{
		Id:        resp.Id,
		BookId:    resp.BookId,
		Inventory: resp.Inventory,
	}, nil
}
