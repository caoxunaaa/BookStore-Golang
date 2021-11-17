package action

import (
	"WebApi/Pb/action"
	"WebApi/Svc"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCommentsByBookContentIdHandler(c *gin.Context) {
	bookContentId, err := strconv.ParseInt(c.Query("bookContentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	//树状结构（评论)
	res, err := Svc.SvcContext.Grpc.ActionGrpc.GetCommentsByBookContentId(ctx, &action.CommentReq{
		BookContentId: bookContentId,
	})
	//树状结构 平铺为 只有父子节点结构（评论)
	tn := func(t *action.CommentsTreeResp) action.CommentsTreeResp {
		var tree = action.CommentsTreeResp{}
		for i := 0; i < len(t.CommentsTree); i++ {
			//组合父节点
			tree.CommentsTree = append(tree.CommentsTree, &action.CommentsNodeResp{
				Comments: t.CommentsTree[i].Comments,
			})
			//组合父节点下所有的节点
			comb(t.CommentsTree[i].CommentsNode, &tree.CommentsTree[i].CommentsNode)
		}
		return tree
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, tn(res))
	}
}

func comb(src []*action.CommentsNodeResp, dest *[]*action.CommentsNodeResp) {
	for i := 0; i < len(src); i++ {
		*dest = append(*dest, &action.CommentsNodeResp{
			Comments: src[i].Comments,
		})
		if src[i].CommentsNode != nil {
			comb(src[i].CommentsNode, dest)
		}
	}

}
