package logic

import (
	"Action/model"
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
	err := l.svcCtx.TrafficStatisticModel.Update(model.TrafficStatistic{
		Id:            in.Id,
		BookId:        in.BookId,
		ChapterNum:    in.ChapterNum,
		TrafficNumber: in.TrafficNumber,
	})
	if err != nil {
		return nil, err
	}
	return &action.Response{Ok: true, Message: "更新阅读量成功"}, nil
}
