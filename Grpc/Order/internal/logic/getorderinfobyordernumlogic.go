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

	return &order.OrderInfoResp{
		Id:          orderInfo.Id,
		BuyerId:     orderInfo.BuyerId,
		OrderNum:    orderInfo.OrderNum,
		OrderTime:   orderInfo.OrderTime.Format("2006-01-02 15:04:05"),
		Cost:        orderInfo.Cost,
		IsPaid:      orderInfo.IsPaid == 1,
		OrderStatus: orderInfo.OrderStatus,
		BookId:      orderInfo.BookId,
	}, nil
}
