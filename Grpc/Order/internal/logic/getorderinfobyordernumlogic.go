package logic

import (
	"context"

	"Order/internal/svc"
	"Order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOrderInfoByOrderNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderInfoByOrderNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoByOrderNumLogic {
	return &GetOrderInfoByOrderNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderInfoByOrderNumLogic) GetOrderInfoByOrderNum(in *order.OrderInfoReq) (*order.OrderInfoResp, error) {
	orderInfo, err := l.svcCtx.OrderInfoModel.FindOneByOrderNum(in.OrderNum)
	if err != nil {
		return nil, err
	}

	orderGoods, err := l.svcCtx.OrderGoodsModel.FindGoodsByOrderNum(in.OrderNum)
	if err != nil {
		return nil, err
	}

	goods := make([]*order.OrderGoods, 0)
	for i, _ := range orderGoods {
		goods = append(goods, &order.OrderGoods{
			Id:       orderGoods[i].Id,
			OrderNum: orderGoods[i].OrderNum,
			BookId:   orderGoods[i].BookId,
			Count:    orderGoods[i].Count,
		})
	}

	return &order.OrderInfoResp{
		Id:          orderInfo.Id,
		BuyerId:     orderInfo.BuyerId,
		OrderNum:    orderInfo.OrderNum,
		CreateTime:  orderInfo.CreateTime.Format("2006-01-02 15:04:05"),
		Cost:        orderInfo.Cost,
		IsPaid:      orderInfo.IsPaid == 1,
		OrderStatus: orderInfo.OrderStatus,
		Goods:       goods,
	}, nil
}
