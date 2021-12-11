package logic

import (
	"context"

	"Action/action"
	"Action/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *action.CommentReq) (*action.Response, error) {
	if err := l.svcCtx.CommentModel.Delete(in.Id); err != nil {
		return nil, err
	}
	return &action.Response{Ok: true, Message: "删除评论成功"}, nil
}
