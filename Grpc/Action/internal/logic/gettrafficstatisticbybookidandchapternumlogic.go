package logic

import (
	"context"

	"Action/action"
	"Action/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTrafficStatisticByBookIdAndChapterNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTrafficStatisticByBookIdAndChapterNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTrafficStatisticByBookIdAndChapterNumLogic {
	return &GetTrafficStatisticByBookIdAndChapterNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTrafficStatisticByBookIdAndChapterNumLogic) GetTrafficStatisticByBookIdAndChapterNum(in *action.TrafficStatisticReq) (*action.TrafficStatisticResp, error) {
	// todo: add your logic here and delete this line

	return &action.TrafficStatisticResp{}, nil
}
