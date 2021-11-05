package logic

import (
	"context"

	"Action/action"
	"Action/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateTrafficStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTrafficStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTrafficStatisticLogic {
	return &CreateTrafficStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTrafficStatisticLogic) CreateTrafficStatistic(in *action.TrafficStatisticReq) (*action.Response, error) {
	// todo: add your logic here and delete this line

	return &action.Response{}, nil
}
