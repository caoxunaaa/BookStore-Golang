package logic

import (
	"Action/model"
	"context"
	"fmt"

	"Action/action"
	"Action/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCommentsByBookContentIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsByBookContentIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsByBookContentIdLogic {
	return &GetCommentsByBookContentIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Comments
func (l *GetCommentsByBookContentIdLogic) GetCommentsByBookContentId(in *action.CommentReq) (*action.CommentsTreeResp, error) {
	comments, err := l.svcCtx.CommentModel.FindCommentsByBookContentId(in.BookContentId)
	if err != nil {
		return nil, err
	}
	fmt.Println("comments", comments)

	f := func(cs []*model.Comment) *action.CommentsTreeResp {
		//树结构
		var res = action.CommentsTreeResp{
			CommentsTree: make([]*action.CommentsNodeResp, 0),
		}
		for i := 0; i < len(cs); i++ {
			if cs[i].ParentId == 0 {
				res.CommentsTree = append(res.CommentsTree, &action.CommentsNodeResp{
					Comments: &action.CommentResp{
						Id:                cs[i].Id,
						ParentId:          cs[i].ParentId,
						BookContentId:     cs[i].BookContentId,
						Comment:           cs[i].Comment,
						CommentByUserId:   cs[i].CommentByUserId,
						CommentByNickname: cs[i].CommentByNickname,
						CommentToUserId:   cs[i].CommentToUserId,
						CommentToNickname: cs[i].CommentToNickname,
					},
				})
			} else {
				node := FindCommentNodeByParentId(res.CommentsTree, cs[i].ParentId)
				if node != nil {
					node.CommentsNode = append(node.CommentsNode, &action.CommentsNodeResp{
						Comments: &action.CommentResp{
							Id:                cs[i].Id,
							ParentId:          cs[i].ParentId,
							BookContentId:     cs[i].BookContentId,
							Comment:           cs[i].Comment,
							CommentByUserId:   cs[i].CommentByUserId,
							CommentByNickname: cs[i].CommentByNickname,
							CommentToUserId:   cs[i].CommentToUserId,
							CommentToNickname: cs[i].CommentToNickname,
						},
					})
				}
			}
		}
		return &res
	}
	return f(comments), nil
}

//找到id对应的节点
func FindCommentNodeByParentId(res []*action.CommentsNodeResp, id int64) *action.CommentsNodeResp {
	for i := 0; i < len(res); i++ {
		if id == res[i].Comments.Id {
			return res[i]
		} else {
			if r := FindCommentNodeByParentId(res[i].CommentsNode, id); r != nil {
				return r
			}
		}
	}
	return nil
}
