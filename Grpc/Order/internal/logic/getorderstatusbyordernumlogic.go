package logic

import (
	"context"

	"Order/internal/svc"
	"Order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOrderStatusByOrderNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderStatusByOrderNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderStatusByOrderNumLogic {
	return &GetOrderStatusByOrderNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderStatusByOrderNumLogic) GetOrderStatusByOrderNum(in *order.OrderInfoReq) (*order.OrderInfoResp, error) {
	// todo: add your logic here and delete this line

	return &order.OrderInfoResp{}, nil
}
