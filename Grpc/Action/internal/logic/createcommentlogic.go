package logic

import (
	"Action/model"
	"context"

	"Action/action"
	"Action/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *action.CommentReq) (*action.Response, error) {
	_, err := l.svcCtx.CommentModel.Insert(model.Comment{
		ParentId:        in.ParentId,
		BookContentId:   in.ParentId,
		Comment:         in.Comment,
		CommentToUserId: in.CommentToUserId,
		CommentByUserId: in.CommentByUserId,
	})
	if err != nil {
		return nil, err
	}

	return &action.Response{Ok: true, Message: "评论成功"}, nil
}
