package logic

import (
	"context"

	"auth-rpc/internal/svc"
	"auth-rpc/pb/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *auth.UserInfoRequest) (*auth.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &auth.UserInfoResponse{}, nil
}
