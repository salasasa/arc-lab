// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package public

import (
	"context"

	"gateway-api/internal/svc"
	"gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIntroLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 简单的自我介绍 (公开)
func NewGetIntroLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIntroLogic {
	return &GetIntroLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIntroLogic) GetIntro() (resp *types.ChatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
