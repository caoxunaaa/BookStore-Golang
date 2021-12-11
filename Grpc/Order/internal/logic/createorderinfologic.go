package logic

import (
	"Order/model"
	"context"
	"fmt"
	"time"

	"Order/internal/svc"
	"Order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateOrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderInfoLogic {
	return &CreateOrderInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderInfoLogic) CreateOrderInfo(in *order.OrderInfoReq) (*order.Response, error) {
	fmt.Println(in)
	_, err := l.svcCtx.OrderInfoModel.Insert(model.OrderInfo{
		BuyerId:     in.BuyerId,
		OrderNum:    in.OrderNum,
		OrderTime:   time.Now(),
		Cost:        in.Cost,
		IsPaid:      0,
		OrderStatus: "待支付",
		BookId:      in.BookId,
	})
	if err != nil {
		return nil, err
	}

	return &order.Response{Ok: true, Message: "成功创建订单"}, nil
}
