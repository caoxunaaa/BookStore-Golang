package logic

import (
	"Order/model"
	"context"
	"time"

	"Order/internal/svc"
	"Order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateOrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderInfoLogic {
	return &UpdateOrderInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderInfoLogic) UpdateOrderInfo(in *order.OrderInfoReq) (*order.Response, error) {
	var isPaid int64
	if in.IsPaid {
		isPaid = 1
	} else {
		isPaid = 0
	}

	err := l.svcCtx.OrderInfoModel.Update(model.OrderInfo{
		Id:          in.Id,
		BuyerId:     in.BuyerId,
		OrderNum:    in.OrderNum,
		CreateTime:  time.Now(),
		Cost:        in.Cost,
		IsPaid:      isPaid,
		OrderStatus: in.OrderStatus,
	})
	if err != nil {
		return nil, err
	}
	return &order.Response{Ok: true, Message: "修改订单状态成功"}, nil
}
