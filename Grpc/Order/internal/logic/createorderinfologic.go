package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &order.Response{}, nil
}
