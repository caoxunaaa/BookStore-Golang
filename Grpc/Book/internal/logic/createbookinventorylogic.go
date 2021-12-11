package logic

import (
	"Book/model"
	"context"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateBookInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBookInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBookInventoryLogic {
	return &CreateBookInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBookInventoryLogic) CreateBookInventory(in *book.BookInventoryReq) (*book.Reply, error) {
	_, err := l.svcCtx.BookInventoryModel.Insert(model.BookInventory{
		BookId:    in.BookId,
		Inventory: in.Inventory,
	})
	if err != nil {
		return nil, err
	}
	return &book.Reply{Ok: true, Message: "创建库存成功"}, nil
}
