package logic

import (
	"context"

	"Action/action"
	"Action/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllTrafficStatisticsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllTrafficStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllTrafficStatisticsLogic {
	return &GetAllTrafficStatisticsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TrafficStatistic
func (l *GetAllTrafficStatisticsLogic) GetAllTrafficStatistics(in *action.Request) (*action.TrafficStatisticsResp, error) {
	// todo: add your logic here and delete this line

	return &action.TrafficStatisticsResp{}, nil
}
