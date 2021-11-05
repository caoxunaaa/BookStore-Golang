package logic

import (
	"Action/model"
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
	reps, err := l.svcCtx.TrafficStatisticModel.FindAll()
	if err != nil {
		return nil, err
	}
	f := func(t []*model.TrafficStatistic) []*action.TrafficStatisticResp {
		var res = make([]*action.TrafficStatisticResp, 0)
		for i := 0; i < len(t); i++ {
			res = append(res, &action.TrafficStatisticResp{
				Id:            t[i].Id,
				BookId:        t[i].BookId,
				ChapterNum:    t[i].ChapterNum,
				TrafficNumber: t[i].TrafficNumber,
			})
		}
		return res
	}

	return &action.TrafficStatisticsResp{TrafficStatistics: f(reps)}, nil
}
