package logic

import (
	"context"

	"Book/book"
	"Book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindOneBookByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneBookByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneBookByIdLogic {
	return &FindOneBookByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOneBookByIdLogic) FindOneBookById(in *book.BookBasicInfoReq) (*book.BookBasicInfoReply, error) {
	rep, err := l.svcCtx.Model.FindOne(in.Id)
	if err != nil {
		return nil, err
	}

	return &book.BookBasicInfoReply{
		Id:          rep.Id,
		Name:        rep.Name,
		Author:      rep.Author,
		Image:       rep.Image,
		StorageTime: rep.StorageTime.Format("2006-01-02"),
	}, nil
}