package logic

import (
	"Book/model"
	"context"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindAllBookInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllBookInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllBookInventoryLogic {
	return &FindAllBookInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  book_inventory
func (l *FindAllBookInventoryLogic) FindAllBookInventory(in *book.BookInventoryReq) (*book.BooksInventoryResp, error) {
	reps, err := l.svcCtx.BookInventoryModel.FindAll()
	if err != nil {
		return nil, err
	}
	f := func(t []*model.BookInventory) []*book.BookInventoryResp {
		var res = make([]*book.BookInventoryResp, 0)
		for i := 0; i < len(t); i++ {
			res = append(res, &book.BookInventoryResp{
				Id:        t[i].Id,
				BookId:    t[i].BookId,
				Inventory: t[i].Inventory,
			})
		}
		return res
	}
	return &book.BooksInventoryResp{BooksInventoryResp: f(reps)}, nil
}
