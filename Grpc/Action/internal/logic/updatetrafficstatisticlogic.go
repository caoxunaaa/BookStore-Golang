package logic

import (
	"context"

	"Action/action"
	"Action/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateTrafficStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTrafficStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTrafficStatisticLogic {
	return &UpdateTrafficStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTrafficStatisticLogic) UpdateTrafficStatistic(in *action.TrafficStatisticReq) (*action.Response, error) {
	// todo: add your logic here and delete this line

	return &action.Response{}, nil
}
