package logic

import (
	"context"

	"Order/internal/svc"
	"Order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteOrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderInfoLogic {
	return &DeleteOrderInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOrderInfoLogic) DeleteOrderInfo(in *order.OrderInfoReq) (*order.Response, error) {
	err := l.svcCtx.OrderInfoModel.Delete(in.Id)
	if err != nil {
		return nil, err
	}
	return &order.Response{Ok: true, Message: "DELETE订单成功"}, nil
}
