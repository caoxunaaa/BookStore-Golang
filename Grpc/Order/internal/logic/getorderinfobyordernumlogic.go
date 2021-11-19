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
	// todo: add your logic here and delete this line

	return &order.OrderInfoResp{}, nil
}
