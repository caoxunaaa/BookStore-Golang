package logic

import (
	"Action/model"
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
	_, err := l.svcCtx.TrafficStatisticModel.Insert(model.TrafficStatistic{
		BookId:        in.BookId,
		ChapterNum:    in.ChapterNum,
		TrafficNumber: in.TrafficNumber,
	})
	if err != nil {
		return nil, err
	}
	return &action.Response{Ok: true, Message: "创建成功"}, nil
}
