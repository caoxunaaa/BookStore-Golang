package logic

import (
	"context"

	"Order/internal/svc"
	"Order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetNotPaidOrderInfoByBuyerIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNotPaidOrderInfoByBuyerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNotPaidOrderInfoByBuyerIdLogic {
	return &GetNotPaidOrderInfoByBuyerIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNotPaidOrderInfoByBuyerIdLogic) GetNotPaidOrderInfoByBuyerId(in *order.OrderInfoReq) (*order.OrderInfoResp, error) {
	orderInfo, err := l.svcCtx.OrderInfoModel.FindNotPaidOrdersByBuyerId(in.BuyerId)
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
