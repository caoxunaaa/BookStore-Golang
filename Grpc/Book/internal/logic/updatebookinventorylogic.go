package logic

import (
	"Book/model"
	"context"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateBookInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBookInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBookInventoryLogic {
	return &UpdateBookInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBookInventoryLogic) UpdateBookInventory(in *book.BookInventoryReq) (*book.Reply, error) {
	err := l.svcCtx.BookInventoryModel.Update(model.BookInventory{
		Id:        in.Id,
		BookId:    in.BookId,
		Inventory: in.Inventory,
	})
	if err != nil {
		return nil, err
	}
	return &book.Reply{Ok: true, Message: "更新库存成功"}, nil
}
