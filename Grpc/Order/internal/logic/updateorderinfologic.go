package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &order.Response{}, nil
}
