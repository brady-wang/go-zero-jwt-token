package logic

import (
	"context"
	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserInfoLogic {
	return GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req types.SearchReq) (*types.SearchReply, error) {

	return &types.SearchReply{}, nil
}
