package logic

import (
	"Order/model"
	"context"
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
	_, err := l.svcCtx.OrderInfoModel.Insert(model.OrderInfo{
		BuyerId:     in.BuyerId,
		OrderNum:    in.OrderNum,
		CreateTime:  time.Now(),
		Cost:        in.Cost,
		IsPaid:      0,
		OrderStatus: "待支付",
	})
	if err != nil {
		return nil, err
	}
	for i, _ := range in.Goods {
		_, err = l.svcCtx.OrderGoodsModel.Insert(model.OrderGoods{
			OrderNum: in.Goods[i].OrderNum,
			BookId:   in.Goods[i].BookId,
			Count:    in.Goods[i].Count,
		})
		if err != nil {
			return nil, err
		}
	}

	return &order.Response{Ok: true, Message: "成功创建订单"}, nil
}
